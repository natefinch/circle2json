package lib

// Room bit definitions
const (

	ROOM_DARK       = 1 << iota // Room is dark.
	ROOM_DEATH                  // Room is a death trap; char ``dies'' (no xp lost).
	ROOM_NOMOB                  // MOBs (monsters) cannot enter room.
	ROOM_INDOORS                // Room is indoors.
	ROOM_PEACEFUL               // Room is peaceful (violence not allowed).
	ROOM_SOUNDPROOF             // Shouts, gossips, etc. won't be heard in room.
	ROOM_NOTRACK                // ``track'' can't find a path through this room.
	ROOM_NOMAGIC                // All magic attempted in this room will fail.
	ROOM_TUNNEL                 // Only one person allowed in room at a time.
	ROOM_PRIVATE                // Cannot teleport in or GOTO if two people here.
	ROOM_GODROOM                // Only LVL_GOD and above allowed to enter.
	ROOM_NOTELEPORT             // You cannot leave this room by any spell
	ROOM_NORELOCATE             // You cannot enter this room by any spell
	ROOM_NOQUIT                 // You cannot quit out from this room
	ROOM_NOFLEE                 // You cannot flee into this room
	ROOM_MAGICDARK              // This room will be dark to mortals, but they can see the room description
	ROOM_BEAMUP                 // If you use a beamer device in your zone, it can only beam people up from rooms that are flagged BEAMUP
	ROOM_FLY                    //  You must be flying to enter this room.
	ROOM_STASIS                 // This gives the room the same effect as the stasis field skill.
)

// LetterBits converts a letter-style bit to the corresponding bit name
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
	'l': "NOTELEPORT",
	'm': "NORELOCATE",
	'n': "NOQUIT",
	'o': "NOFLEE",
	'p': "MAGICDARK",
	'q': "BEAMUP",
	'r': "FLY",
	's': "STASIS",
}


var RoomBits = map[int]string{
	ROOM_DARK:       "DARK",
	ROOM_DEATH:      "DEATH",
	ROOM_NOMOB:      "NOMOB",
	ROOM_INDOORS:    "INDOORS",
	ROOM_PEACEFUL:   "PEACEFUL",
	ROOM_SOUNDPROOF: "SOUNDPROOF",
	ROOM_NOTRACK:    "NOTRACK",
	ROOM_NOMAGIC:    "NOMAGIC",
	ROOM_TUNNEL:     "TUNNEL",
	ROOM_PRIVATE:    "PRIVATE",
	ROOM_GODROOM:    "GODROOM",
	ROOM_NOTELEPORT: "NOTELEPORT",
	ROOM_NORELOCATE: "NORELOCATE",
	ROOM_NOQUIT:     "NOQUIT",
	ROOM_NOFLEE:     "NOFLEE",
	ROOM_MAGICDARK:  "MAGICDARK",
	ROOM_BEAMUP:     "BEAMUP",
	ROOM_FLY:        "FLY",
	ROOM_STASIS:     "STASIS",
}

// RoomBitsToNames converts a room's bitvector into a list of bit names
func RoomBitsToNames(vector string) ([]string, error) {
	return BitsToNames(vector, ROOM_STASIS, RoomBits, RoomChars)
}
