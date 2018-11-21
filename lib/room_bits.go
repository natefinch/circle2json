package lib

import (
	"fmt"
	"strconv"
)

// Room bit definitions
const (
	DARK       = 1 << iota // Room is dark.
	DEATH                  // Room is a death trap; char ``dies'' (no xp lost).
	NOMOB                  // MOBs (monsters) cannot enter room.
	INDOORS                // Room is indoors.
	PEACEFUL               // Room is peaceful (violence not allowed).
	SOUNDPROOF             // Shouts, gossips, etc. won't be heard in room.
	NOTRACK                // ``track'' can't find a path through this room.
	NOMAGIC                // All magic attempted in this room will fail.
	TUNNEL                 // Only one person allowed in room at a time.
	PRIVATE                // Cannot teleport in or GOTO if two people here.
	GODROOM                // Only LVL_GOD and above allowed to enter.
	NOTELEPORT             // You cannot leave this room by any spell
	NORELOCATE             // You cannot enter this room by any spell
	NOQUIT                 // You cannot quit out from this room
	NOFLEE                 // You cannot flee into this room
	MAGICDARK              // This room will be dark to mortals, but they can see the room description
	BEAMUP                 // If you use a beamer device in your zone, it can only beam people up from rooms that are flagged BEAMUP
	FLY                    //  You must be flying to enter this room.
	STASIS                 // This gives the room the same effect as the stasis field skill.
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
	'l': "NOTELEPORT",
	'm': "NORELOCATE",
	'n': "NOQUIT",
	'o': "NOFLEE",
	'p': "MAGICDARK",
	'q': "BEAMUP",
	'r': "FLY",
	's': "STASIS",
}

var BitNames = map[int]string{
	DARK:       "DARK",
	DEATH:      "DEATH",
	NOMOB:      "NOMOB",
	INDOORS:    "INDOORS",
	PEACEFUL:   "PEACEFUL",
	SOUNDPROOF: "SOUNDPROOF",
	NOTRACK:    "NOTRACK",
	NOMAGIC:    "NOMAGIC",
	TUNNEL:     "TUNNEL",
	PRIVATE:    "PRIVATE",
	GODROOM:    "GODROOM",
	NOTELEPORT: "NOTELEPORT",
	NORELOCATE: "NORELOCATE",
	NOQUIT:     "NOQUIT",
	NOFLEE:     "NOFLEE",
	MAGICDARK:  "MAGICDARK",
	BEAMUP:     "BEAMUP",
	FLY:        "FLY",
	STASIS:     "STASIS",
}

// BitVectorToNames converts a room's bitvector into a list of bit names
func BitVectorToNames(vector string) ([]string, error) {
	values := []string{}
	if num, err := strconv.Atoi(vector); err == nil {
		// number-style bitvector
		for bit := DARK; bit <= STASIS; bit = bit << 1 {
			if num&bit != 0 {
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
