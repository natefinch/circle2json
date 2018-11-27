package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/natefinch/circle2json/lib"
)

func main() {
	var to, from, pattern string
	var mode string
	flag.StringVar(&from, "from", ".", "specifies the input directory")
	flag.StringVar(&to, "to", "./json", "specifies the output directory")
	flag.StringVar(&pattern, "pattern", "*.wld", "specifies the glob pattern used to find files")
	flag.StringVar(&mode, "mode", "room", "mob, zone, or room (defaults pattern to *.mob, *.zon *.wld, respectively)")
	flag.Usage = func() {
		fmt.Print("circle2json converts Ashes2Ashes world files into json files.\n\n")
		fmt.Print("usage: circle2json [options]\n\n")
		flag.PrintDefaults()
		fmt.Print("  -help\n        show this help\n")
	}
	flag.Parse()

	log.SetFlags(0)
	switch mode {
	case "zone", "zones":
		if pattern == "*.wld" {
			pattern = "*.zon"
		}
		if err := lib.ConvertZones(to, from, pattern); err != nil {
			log.Fatal(err)
		}
	case "room", "rooms":
		if err := lib.ConvertRooms(to, from, pattern); err != nil {
			log.Fatal(err)
		}
	case "mob", "mobs":
		if pattern == "*.wld" {
			pattern = "*.mob"
		}
		if err := lib.ConvertMobs(to, from, pattern); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown mode: %v", mode)
	}
	log.Println("success!")
}
