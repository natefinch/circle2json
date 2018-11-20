package lib

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
