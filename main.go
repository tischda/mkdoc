package main

import (
	"flag"
	"fmt"
	"log"
)

// http://technosophos.com/2014/06/11/compile-time-string-in-go.html
// go build -ldflags "-x main.version $(git describe --tags)"
var version string

// command line flags
var showVersion bool
var withRenumber bool
var imgCheckDir string

func init() {
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
	flag.BoolVar(&withRenumber, "renumber", false, "renumber markdown source files")
	flag.StringVar(&imgCheckDir, "check", "", "check image directory for orphans")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if showVersion {
		fmt.Println("mkdoc version", version)
	} else {
		if imgCheckDir != "" {
			listOrphans()
		}
		if withRenumber {
			renumberFiles()
		}
		runPandoc()
	}
}
