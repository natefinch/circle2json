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

// ConvertMobs converts all the CircleMUD mob files in the from directory
// that match the pattern to json files in the to directory.
func ConvertMobs(to, from, pattern string) (err error) {
	if err := os.MkdirAll(to, 0700); err != nil {
		return fmt.Errorf("couldn't create output directory: %v", err)
	}
	files, err := filepath.Glob(filepath.Join(from, pattern))
	if err != nil {
		return err
	}
	for _, name := range files {
		mobs, err := ParseMobFile(name)
		if err != nil {
			return err
		}
		b, err := json.MarshalIndent(map[string]interface{}{"mobs": mobs}, "", "    ")
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

// ParseMobFile parses the given CircleMUD mob file.
func ParseMobFile(filename string) (_ []*Mob, err error) {
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
	mobs := []*Mob{}
	if err := scanner.MustScan(); err != nil {
		return nil, err
	}
	for {
		if strings.TrimSpace(scanner.Text()) == "$" {
			return mobs, nil
		}
		mob, err := scanMob(scanner)
		if err != nil {
			return nil, err
		}
		mobs = append(mobs, mob)
	}
}

func scanMob(scanner *fileScanner) (*Mob, error) {
	number := strings.TrimSpace(scanner.Text())
	if !strings.HasPrefix(number, "#") {
		return nil, fmt.Errorf("mob number must start with #, but found: %q", number)
	}
	num, err := strconv.Atoi(number[1:])
	if err != nil {
		return nil, fmt.Errorf("mob number %q not a number: %v", number[1:], err)
	}
	m := Mob{Number: num}

	d, err := scanner.ScanUntil("~")
	if err != nil {
		return nil, err
	}
	m.Aliases = strings.Split(d, " ")

	d, err = scanner.ScanUntil("~")
	if err != nil {
		return nil, err
	}
	m.ShortDesc = d
	d, err = scanner.ScanUntil("~")
	if err != nil {
		return nil, err
	}
	m.LongDesc = d
	d, err = scanner.ScanUntil("~")
	if err != nil {
		return nil, err
	}
	m.DetailedDesc = d
	if err := scanner.MustScan(); err != nil {
		return nil, err
	}
	fields := strings.Fields(scanner.Text())
	if len(fields) != 4 {
		return nil, fmt.Errorf("expected mob metadata to be <action_bits> <affection_bits> <alignment> <type>, but got %q", scanner.Text())
	}
	actions, err := MobActionsToNames(fields[0])
	if err != nil {
		return nil, err
	}
	m.Actions = actions

	affections, err := MobAffectionsToNames(fields[1])
	if err != nil {
		return nil, err
	}
	m.Affections = affections

	alignment, err := strconv.Atoi(fields[2])
	if err != nil {
		return nil, err
	}
	m.Alignment = alignment

	mobtype := fields[3]
	switch mobtype {
	case "S", "E", "W", "W1", "W2", "W3":
		// ok
	default:
		return nil, fmt.Errorf("expected mob type to be S, W, or E, but was %v", mobtype)
	}
	// I like how there's 3 different "types" but then they all have the same first 3 lines....
	if err := scanner.MustScan(); err != nil {
		return nil, err
	}
	fields = strings.Fields(scanner.Text())
	if len(fields) != 5 {
		return nil, fmt.Errorf("expected mob metadata to be <level> <thac0> <armor class> <max hit points> <bare hand damage>, but got %q", scanner.Text())
	}
	level, err := strconv.Atoi(fields[0])
	if err != nil {
		return nil, fmt.Errorf("invalid level: %v", err)
	}
	m.Level = level

	thac0, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, fmt.Errorf("invalid thac0: %v", err)
	}
	m.THAC0 = thac0

	ac, err := strconv.Atoi(fields[2])
	if err != nil {
		return nil, fmt.Errorf("invalid AC: %v", err)
	}
	m.AC = ac

	m.HP = fields[3]
	m.Damage = fields[4]

	if err := scanner.MustScan(); err != nil {
		return nil, err
	}
	fields = strings.Fields(scanner.Text())
	if len(fields) != 2 {
		return nil, fmt.Errorf("expected mob metadata to be <gold> <experience points>, but got %q", scanner.Text())
	}
	gold, err := strconv.Atoi(fields[0])
	if err != nil {
		return nil, err
	}
	m.Gold = gold

	xp, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, err
	}
	m.XP = xp

	if err := scanner.MustScan(); err != nil {
		return nil, err
	}
	fields = strings.Fields(scanner.Text())
	if len(fields) != 3 {
		return nil, fmt.Errorf("expected mob metadata to be <load position> <default position> <sex>, but got %q", scanner.Text())
	}
	pos, ok := PositionNames[fields[0]]
	if !ok {
		return nil, fmt.Errorf("unknown position: %s", fields[0])
	}
	m.LoadPosition = pos
	pos, ok = PositionNames[fields[1]]
	if !ok {
		return nil, fmt.Errorf("unknown position: %s", fields[1])
	}
	m.DefaultPosition = pos

	gender, ok := genders[fields[2]]
	if !ok {
		return nil, fmt.Errorf("unknown gender: %s", fields[2])
	}
	m.Gender = gender

	// for now just discard anything else
	if _, err := scanner.ScanUntilPrefix([]string{"#", "$"}); err != nil {
		return nil, err
	}

	return &m, nil
}

type Mob struct {
	Number          int
	Aliases         []string
	ShortDesc       string
	LongDesc        string
	DetailedDesc    string
	Actions         []string
	Affections      []string
	Alignment       int
	Level           int
	THAC0           int
	AC              int
	HP              string // xdy+z
	Damage          string // xdy+z
	Gold            int
	XP              int
	LoadPosition    string
	DefaultPosition string
	Gender          string
}

var PositionNames = map[string]string{
	"0": "POSITION_DEAD",
	"1": "POSITION_MORTALLYW",
	"2": "POSITION_INCAP",
	"3": "POSITION_STUNNED",
	"4": "POSITION_SLEEPING",
	"5": "POSITION_RESTING",
	"6": "POSITION_SITTING",
	"7": "POSITION_FIGHTING",
	"8": "POSITION_STANDING",
}

var genders = map[string]string{
	"0": "Neutral",
	"1": "Male",
	"2": "Female",
}
