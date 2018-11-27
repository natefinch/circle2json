# circle2json

```
circle2json converts CircleMUD world (room) files into json files.

usage: circle2json [options]

  -from string
        specifies the input directory (default ".")
  -mode string
        mob, zone, or room (defaults pattern to *.mob, *.zon *.wld, respectively) (default "room")
  -pattern string
        specifies the glob pattern used to find files (default "*.wld")
  -to string
        specifies the output directory (default "./json")
  -help
        show this help
```

The structure of the output is given here (as defined in Go):

```go

// Room is a representation of a room in a MUD.
type Room struct {
	Number      int         `json:"number"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Bits        []string    `json:"bits"`
	Sector      string      `json:"sector"`
	Exits       []Exit      `json:"exits"`
	Extras      []ExtraDesc `json:"extra_descs"`
}

// Exit represents a way you may move out of a room.
type Exit struct {
	Direction   string   `json:"direction"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
	DoorFlag    string   `json:"door_flag"`
	KeyNumber   int      `json:"key_number"`
	Destination int      `json:"destination"`
}

// ExtraDesc represents other things you can look at in the room.
type ExtraDesc struct {
	Keywords    []string `json:"keywords"`
	Description string   `json:"description"`
}

// Room bits define different flags of a room.
const (
	DARK        = "DARK"        // Room is dark.
	DEATH       = "DEATH"       // Room is a death trap; char dies (no xp lost).
	NOMOB       = "NOMOB"       // MOBs (monsters) cannot enter room.
	INDOORS     = "INDOORS"     // Room is indoors.
	PEACEFUL    = "PEACEFUL"    // Room is peaceful (violence not allowed).
	SOUNDPROOF  = "SOUNDPROOF"  // Shouts, gossips, etc. won't be heard in room.
	NOTRACK     = "NOTRACK"     // track can't find a path through this room.
	NOMAGIC     = "NOMAGIC"     // All magic attempted in this room will fail.
	TUNNEL      = "TUNNEL"      // Only one person allowed in room at a time.
	PRIVATE     = "PRIVATE"     // Cannot teleport in or GOTO if two people here.
	GODROOM     = "GODROOM"     // Only LVL_GOD and above allowed to enter.
	HOUSE       = "HOUSE"       // Reserved for internal use.  Do not set.
	HOUSE_CRASH = "HOUSE_CRASH" // Reserved for internal use.  Do not set.
	ATRIUM      = "ATRIUM"      // Reserved for internal use.  Do not set.
	OLC         = "OLC"         // Reserved for internal use.  Do not set.
	BFS_MARK    = "BFS_MARK"    // Reserved for internal use.  Do not set.
)



// SectorType is a conversion between CircleMUD's sector type number and a human-readable string.
var SectorType = map[string]string{
	"0": "INSIDE",       // Indoors (small number of move points needed).
	"1": "CITY",         // The streets of a city.
	"2": "FIELD",        // An open field.
	"3": "FOREST",       // A dense forest.
	"4": "HILLS",        // Low foothills.
	"5": "MOUNTAIN",     // Steep mountain regions.
	"6": "WATER_SWIM",   // Water (swimmable).
	"7": "WATER_NOSWIM", // Unswimmable water - boat required for passage.
	"8": "UNDERWATER",   // Underwater.
	"9": "FLYING",       // Wheee!
}

// ExitDir is the conversion between CircleMUD's direction number and a human-readable string.
var ExitDir = map[string]string{
	"0": "North",
	"1": "East",
	"2": "South",
	"3": "West",
	"4": "Up",
	"5": "Down",
}

// DoorFlags is the conversion between CircleMUD's door flags and a human-readable string.
var DoorFlags = map[string]string{
	"0": "NONE",
	"1": "NORMAL",
	"2": "PICKPROOF",
}
```

....that's basically it.  Then I presume anyone can convert from json to an in-memory copy in their own language of choice.
