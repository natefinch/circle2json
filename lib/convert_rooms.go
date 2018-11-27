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

// ConvertRooms converts all the CircleMUD world (room) files in the from directory
// that match the pattern to json files in the to directory.
func ConvertRooms(to, from, pattern string) (err error) {
	if err := os.MkdirAll(to, 0700); err != nil {
		return fmt.Errorf("couldn't create output directory: %v", err)
	}
	files, err := filepath.Glob(filepath.Join(from, pattern))
	if err != nil {
		return err
	}
	for _, name := range files {
		r, err := ParseWldFile(name)
		if err != nil {
			return err
		}
		output := struct {
			Rooms []Room `json:"rooms"`
		}{
			Rooms: r,
		}
		b, err := json.MarshalIndent(output, "", "    ")
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

// Room is a representation of a room in a MUD.
type Room struct {
	Number      int         `json:"ID"`
	Zone        int         `json:"Zone"`
	Name        string      `json:"Name"`
	Description string      `json:"Description"`
	Bits        []string    `json:"Bits"`
	Sector      string      `json:"Sector"`
	Exits       []Exit      `json:"Exits"`
	Extras      []ExtraDesc `json:"ExtraDescs"`
}

// Exit represents a way you may move out of a room.
type Exit struct {
	Direction   string   `json:"Direction"`
	Description string   `json:"Description"`
	Keywords    []string `json:"Keywords"`
	DoorFlags   []string `json:"DoorFlags"`
	KeyNumber   int      `json:"KeyID"`
	Destination int      `json:"Destination"`
}

// ExtraDesc represents other things you can look at in the room.
type ExtraDesc struct {
	Keywords    []string `json:"Keywords"`
	Description string   `json:"Description"`
}

// ParseWldFile parses the given CircleMUD wld file.
func ParseWldFile(filename string) (rooms []Room, err error) {
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
	for {
		if !scanner.Scan() {
			if err = scanner.Err(); err != nil {
				return nil, err
			}
			// end of file, that's ok. Technically you're supposed to end the
			// file with $, but it doesn't really seem to be necessary.
			return rooms, nil
		}
		if strings.TrimSpace(scanner.Text()) == "$" {
			return rooms, nil
		}
		room, err := scanRoom(scanner)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, *room)
	}
}

func scanRoom(scanner *fileScanner) (*Room, error) {
	number := strings.TrimSpace(scanner.Text())
	if !strings.HasPrefix(number, "#") {
		return nil, fmt.Errorf("room number must start with #, but found: %q", number)
	}
	num, err := strconv.Atoi(number[1:])
	if err != nil {
		return nil, fmt.Errorf("room number %q not a number: %v", number[1:], err)
	}
	r := Room{Number: num}
	name, err := scanner.ScanUntil("~")
	if err != nil {
		return nil, err
	}
	r.Name = name
	desc, err := scanner.ScanUntil("~")
	if err != nil {
		return nil, err
	}
	r.Description = desc
	if err := scanner.MustScan(); err != nil {
		return nil, err
	}
	fields := strings.Fields(scanner.Text())
	if len(fields) != 3 {
		return nil, fmt.Errorf("expected room metadata to be <zone#> <bitvector> <sector>, but got %q", scanner.Text())
	}

	zone, err := strconv.Atoi(fields[0])
	if err != nil {
		return nil, fmt.Errorf("invalid zone number: %q", fields[0])
	}
	r.Zone = zone

	bits, err := RoomBitsToNames(fields[1])
	if err != nil {
		return nil, err
	}
	r.Bits = bits
	sector, ok := SectorType[fields[2]]
	if !ok {
		return nil, fmt.Errorf("unknown room sector type: %q", fields[2])
	}
	r.Sector = sector

	// ignore damage in the room for now.
	if err := scanner.MustScan(); err != nil {
		return nil, err
	}

	for {
		// optional stuff
		if err := scanner.MustScan(); err != nil {
			return nil, err
		}
		s := strings.TrimSpace(scanner.Text())
		switch {
		case s == "S":
			// end of room
			return &r, nil
		case strings.HasPrefix(s, "D"):
			dir, err := scanDir(scanner)
			if err != nil {
				return nil, err
			}
			r.Exits = append(r.Exits, *dir)
		case s == "E":
			ex, err := scanExtra(scanner)
			if err != nil {
				return nil, err
			}
			r.Extras = append(r.Extras, *ex)
		default:
			return nil, fmt.Errorf("unexpected token in room definition: %q", s)
		}
	}
}

func scanDir(scanner *fileScanner) (*Exit, error) {
	// previous code checked that the first character was a D so we can ignore that.
	s := strings.TrimSpace(scanner.Text()[1:])
	dir, ok := ExitDir[s]
	if !ok {
		return nil, fmt.Errorf("unknown exit direction %q", s)
	}
	ex := &Exit{
		Direction: dir,
	}
	desc, err := scanner.ScanUntil("~")
	if err != nil {
		return nil, err
	}
	ex.Description = desc
	keywords, err := scanner.ScanUntil("~")
	if err != nil {
		return nil, err
	}
	ex.Keywords = strings.Fields(keywords)
	if err := scanner.MustScan(); err != nil {
		return nil, err
	}
	fields := strings.Fields(scanner.Text())
	if len(fields) != 3 {
		return nil, fmt.Errorf("expected direction fields to be <door_flag> <key_number> <room_linked> but got %q", scanner.Text())
	}
	flags, err := doorFlags(fields[0])
	if err != nil {
		return nil, err
	}
	ex.DoorFlags = flags
	num, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, fmt.Errorf("invalid key number: %q", fields[1])
	}
	ex.KeyNumber = num
	room, err := strconv.Atoi(fields[2])
	if err != nil {
		return nil, fmt.Errorf("invalid direction target room number: %q", fields[2])
	}
	ex.Destination = room
	return ex, nil
}

func scanExtra(scanner *fileScanner) (*ExtraDesc, error) {
	s, err := scanner.ScanUntil("~")
	if err != nil {
		return nil, err
	}
	keywords := strings.Fields(s)
	ex := &ExtraDesc{
		Keywords: keywords,
	}
	desc, err := scanner.ScanUntil("~")
	if err != nil {
		return nil, err
	}
	ex.Description = desc
	return ex, nil
}
