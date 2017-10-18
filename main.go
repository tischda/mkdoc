package main

import (
	"flag"
	"fmt"
	"log"
)

var version string

// command line flags
var showVersion bool
var noop bool
var withRenumber bool
var imgageDir string

func init() {
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
	flag.BoolVar(&noop, "noop", false, "don't execute pandoc (show options)")
	flag.BoolVar(&withRenumber, "renumber", false, "renumber markdown source files")
	flag.StringVar(&imgageDir, "check", "", "check image directory for orphans")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if showVersion {
		fmt.Println("mkdoc version", version)
	} else {
		if imgageDir != "" {
			listOrphans()
		}
		if withRenumber {
			renumberFiles()
		}
		runPandoc()
	}
}
