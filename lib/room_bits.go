package lib

import (
	"fmt"
	"strconv"
)

const (
	DARK        = 1     // Room is dark.
	DEATH       = 2     // Room is a death trap; char ``dies'' (no xp lost).
	NOMOB       = 4     // MOBs (monsters) cannot enter room.
	INDOORS     = 8     // Room is indoors.
	PEACEFUL    = 16    // Room is peaceful (violence not allowed).
	SOUNDPROOF  = 32    // Shouts, gossips, etc. won't be heard in room.
	NOTRACK     = 64    // ``track'' can't find a path through this room.
	NOMAGIC     = 128   // All magic attempted in this room will fail.
	TUNNEL      = 256   // Only one person allowed in room at a time.
	PRIVATE     = 512   // Cannot teleport in or GOTO if two people here.
	GODROOM     = 1024  // Only LVL_GOD and above allowed to enter.
	HOUSE       = 2048  // Reserved for internal use.  Do not set.
	HOUSE_CRASH = 4096  // Reserved for internal use.  Do not set.
	ATRIUM      = 8192  // Reserved for internal use.  Do not set.
	OLC         = 16384 // Reserved for internal use.  Do not set.
	BFS_MARK    = 32768 // Reserved for internal use.  Do not set.
)

// LetterBits converts a letter-style bit to the corresponding bit name
var LetterBits = map[rune]string{
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

var BitNames = map[int]string{
	DARK:        "DARK",
	DEATH:       "DEATH",
	NOMOB:       "NOMOB",
	INDOORS:     "INDOORS",
	PEACEFUL:    "PEACEFUL",
	SOUNDPROOF:  "SOUNDPROOF",
	NOTRACK:     "NOTRACK",
	NOMAGIC:     "NOMAGIC",
	TUNNEL:      "TUNNEL",
	PRIVATE:     "PRIVATE",
	GODROOM:     "GODROOM",
	HOUSE:       "HOUSE",
	HOUSE_CRASH: "HOUSE_CRASH",
	ATRIUM:      "ATRIUM",
	OLC:         "OLC",
	BFS_MARK:    "BFS_MARK",
}

// BitVectorToNames converts a room's bitvector into a list of bit names
func BitVectorToNames(vector string) ([]string, error) {
	values := []string{}
	if num, err := strconv.Atoi(vector); err == nil {
		// number-style bitvector
		for bit := DARK; bit <= BFS_MARK; bit = bit << 1 {
			if num&bit == 1 {
				values = append(values, BitNames[bit])
			}
		}
		return values, nil
	}
	for _, r := range []rune(vector) {
		s, ok := LetterBits[r]
		if !ok {
			return nil, fmt.Errorf("unknown bit vector letter: %v", r)
		}
		values = append(values, s)
	}
	return values, nil
}
