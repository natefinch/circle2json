package lib

// Room bit definitions
const (
	ROOM_DARK        = 1 << iota // Room is dark.
	ROOM_DEATH                   // Room is a death trap; char ``dies'' (no xp lost).
	ROOM_NOMOB                   // MOBs (monsters) cannot enter room.
	ROOM_INDOORS                 // Room is indoors.
	ROOM_PEACEFUL                // Room is peaceful (violence not allowed).
	ROOM_SOUNDPROOF              // Shouts, gossips, etc. won't be heard in room.
	ROOM_NOTRACK                 // ``track'' can't find a path through this room.
	ROOM_NOMAGIC                 // All magic attempted in this room will fail.
	ROOM_TUNNEL                  // Only one person allowed in room at a time.
	ROOM_PRIVATE                 // Cannot teleport in or GOTO if two people here.
	ROOM_GODROOM                 // Only LVL_GOD and above allowed to enter.
	ROOM_HOUSE                   // Reserved for internal use.  Do not set.
	ROOM_HOUSE_CRASH             // Reserved for internal use.  Do not set.
	ROOM_ATRIUM                  // Reserved for internal use.  Do not set.
	ROOM_OLC                     // Reserved for internal use.  Do not set.
	ROOM_BFS_MARK                // Reserved for internal use.  Do not set.
)

// RoomChars converts a letter-style bit to the corresponding bit name
var RoomChars = map[rune]string{
	'a': "DARK",
	'b': "DEATH",
	'c': "NOMOB",
	'd': "INDOORS",
	'e': "PEACEFUL",
	'f': "SOUNDPROOF",
	'g': "NOTRACK",
	'h': "NOMAGIC",
	'i': "TUNNEL",
	'j': "PRIVATE",
	'k': "GODROOM",
	'l': "HOUSE",
	'm': "HOUSE_CRASH",
	'n': "ATRIUM",
	'o': "OLC",
	'p': "BFS_MARK",
}

var RoomBits = map[int]string{
	ROOM_DARK:        "DARK",
	ROOM_DEATH:       "DEATH",
	ROOM_NOMOB:       "NOMOB",
	ROOM_INDOORS:     "INDOORS",
	ROOM_PEACEFUL:    "PEACEFUL",
	ROOM_SOUNDPROOF:  "SOUNDPROOF",
	ROOM_NOTRACK:     "NOTRACK",
	ROOM_NOMAGIC:     "NOMAGIC",
	ROOM_TUNNEL:      "TUNNEL",
	ROOM_PRIVATE:     "PRIVATE",
	ROOM_GODROOM:     "GODROOM",
	ROOM_HOUSE:       "HOUSE",
	ROOM_HOUSE_CRASH: "HOUSE_CRASH",
	ROOM_ATRIUM:      "ATRIUM",
	ROOM_OLC:         "OLC",
	ROOM_BFS_MARK:    "BFS_MARK",
}

// BitVectorToNames converts a room's bitvector into a list of bit names
func BitVectorToNames(vector string) ([]string, error) {
	return BitsToNames(vector, ROOM_BFS_MARK, RoomBits, RoomChars)
}
