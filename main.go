package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/natefinch/circle2json/lib"
)

func main() {
	var to, from, pattern string
	flag.StringVar(&from, "from", ".", "specifies the input directory")
	flag.StringVar(&to, "to", "./json", "specifies the output directory")
	flag.StringVar(&pattern, "pattern", "*.wld", "specifies the glob pattern used to find files")
	flag.Usage = func() {
		fmt.Print("circle2json converts CircleMUD world (room) files into json files.\n\n")
		fmt.Print("usage: circle2json [options]\n\n")
		flag.PrintDefaults()
		fmt.Print("  -help\n        show this help\n")
	}
	flag.Parse()

	if err := lib.Convert(to, from, pattern); err != nil {
		log.Fatal(err)
	}
	log.Println("success!")
}

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
