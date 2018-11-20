package main

import (
	"flag"
	"log"

	"github.com/natefinch/circle2json/lib"
)

func main() {
	var to, from, pattern string
	flag.StringVar(&from, "from", ".", "specifies the input directory")
	flag.StringVar(&to, "to", "./json", "specifies the output directory")
	flag.StringVar(&pattern, "pattern", "*.wld", "specifies the glob pattern used to find files")
	flag.Parse()

	if err := lib.Convert(to, from, pattern); err != nil {
		log.Fatal(err)
	}
	log.Println("success!")
}
