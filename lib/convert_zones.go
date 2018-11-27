package lib

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ConvertZones converts all the CircleMUD zone files in the from directory
// that match the pattern to json files in the to directory.
func ConvertZones(to, from, pattern string) (err error) {
	if err := os.MkdirAll(to, 0700); err != nil {
		return fmt.Errorf("couldn't create output directory: %v", err)
	}
	files, err := filepath.Glob(filepath.Join(from, pattern))
	if err != nil {
		return err
	}
	for _, name := range files {
		zone, err := ParseZoneFile(name)
		if err != nil {
			return err
		}
		b, err := json.MarshalIndent(zone, "", "    ")
		if err != nil {
			return fmt.Errorf("failed to convert %q to json: %v", name, err)
		}
		n := filepath.Base(name)
		ext := filepath.Ext(n)
		n = n[:len(n)-len(ext)] + ".json"
		n = filepath.Join(to, n)

		if err := ioutil.WriteFile(n, b, 0600); err != nil {
			return err
		}
	}
	return nil
}

// ParseZoneFile parses the given CircleMUD zone file.
func ParseZoneFile(filename string) (zone *Zone, err error) {
	// need this because scan can panic if you send it too much stuff
	defer func() {
		panicErr := recover()
		if panicErr == nil {
			return
		}
		if e, ok := panicErr.(error); ok {
			err = e
			return
		}
		err = fmt.Errorf("%v", panicErr)
	}()

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	line := 0
	scanner := &fileScanner{
		line:    &line,
		Scanner: bufio.NewScanner(f),
	}
	defer func() {
		if err != nil {
			// add filename and line number to error
			err = fmt.Errorf("%s:%v - %s", filename, line, err)
		}
	}()
	if err := scanner.MustScan(); err != nil {
		return nil, err
	}
	zone, err = scanZone(scanner)
	if err != nil {
		return nil, err
	}
	if err := scanner.MustScan(); err != nil {
		return nil, err
	}
	if strings.TrimSpace(scanner.Text()) != "$" {
		return nil, fmt.Errorf("unexpected data at end of zone file: %q", scanner.Text())
	}
	return zone, nil
}

type Zone struct {
	Number       int    `json:"ID"`
	Name         string `json:"Name"`
	BottomNumber int    `json:"BottomNumber"`
	TopNumber    int    `json:"TopNumber"`
	LifespanMins int    `json:"Lifespan"`
	ResetMode    string `json:"ResetMode"`
	Closed       bool   `json:"Closed"`
}

const (
	RESET_NEVER  = "RESET_NEVER"
	RESET_EMPTY  = "RESET_EMPTY"
	RESET_ALWAYS = "RESET_ALWAYS"
)

var resetMap = map[string]string{
	"0": RESET_NEVER,
	"1": RESET_EMPTY,
	"2": RESET_ALWAYS,
}

func scanZone(scanner *fileScanner) (*Zone, error) {
	number := strings.TrimSpace(scanner.Text())
	if !strings.HasPrefix(number, "#") {
		return nil, fmt.Errorf("zone number must start with #, but found: %q", number)
	}
	num, err := strconv.Atoi(number[1:])
	if err != nil {
		return nil, fmt.Errorf("zone number %q not a number: %v", number[1:], err)
	}
	z := Zone{Number: num}
	name, err := scanner.ScanUntil("~")
	if err != nil {
		return nil, err
	}
	z.Name = name
	if err := scanner.MustScan(); err != nil {
		return nil, err
	}
	fields := strings.Fields(scanner.Text())
	if len(fields) != 5 {
		return nil, fmt.Errorf("expected zone metadata to be <bottom_room#> <top_room#> <lifespan> <reset_mode> <closed>, but got %q", scanner.Text())
	}

	bottomRoom, err := strconv.Atoi(fields[0])
	if err != nil {
		return nil, fmt.Errorf("invalid bottom room number: %q", fields[0])
	}
	z.BottomNumber = bottomRoom

	topRoom, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, fmt.Errorf("invalid top room number: %q", fields[1])
	}
	z.TopNumber = topRoom
	lifespan, err := strconv.Atoi(fields[2])
	if err != nil {
		return nil, fmt.Errorf("invalid lifespan number: %q", fields[2])
	}
	z.LifespanMins = lifespan
	mode, ok := resetMap[fields[3]]
	if !ok {
		return nil, fmt.Errorf("unknown reset mode: %q", fields[3])
	}
	z.ResetMode = mode
	switch fields[4] {
	case "0":
		z.Closed = false
	case "1":
		z.Closed = true
	default:
		return nil, fmt.Errorf("unexpected value for `closed`: %q", fields[4])
	}

	// skip all mobs and zone commands for now.
	_, err = scanner.ScanUntil("S")
	if err != nil {
		return nil, err
	}
	return &z, nil
}
