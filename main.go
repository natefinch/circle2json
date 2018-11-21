package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/natefinch/circle2json/lib"
)

func main() {
	var to, from, pattern string
	var zone bool
	flag.StringVar(&from, "from", ".", "specifies the input directory")
	flag.StringVar(&to, "to", "./json", "specifies the output directory")
	flag.StringVar(&pattern, "pattern", "*.wld", "specifies the glob pattern used to find files")
	flag.BoolVar(&zone, "zone", false, "parse zone files instead of room files (makes pattern *.zon)")
	flag.Usage = func() {
		fmt.Print("circle2json converts CircleMUD world (room) files into json files.\n\n")
		fmt.Print("usage: circle2json [options]\n\n")
		flag.PrintDefaults()
		fmt.Print("  -help\n        show this help\n")
	}
	flag.Parse()

	log.SetFlags(0)
	if zone {
		if pattern == "*.wld" {
			pattern = "*.zon"
		}
		if err := lib.ConvertZones(to, from, pattern); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := lib.ConvertRooms(to, from, pattern); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("success!")
}
