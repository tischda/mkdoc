package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// https://goreleaser.com/cookbooks/using-main.version/
var (
	name    string
	version string
	date    string
	commit  string
)

// flags
type Config struct {
	imagePath string
	noop      bool
	renumber  bool
	help      bool
	version   bool
}

func initFlags() *Config {
	cfg := &Config{}
	flag.StringVar(&cfg.imagePath, "c", "", "")
	flag.StringVar(&cfg.imagePath, "check", "", "check image directory for orphans")
	flag.BoolVar(&cfg.noop, "n", false, "")
	flag.BoolVar(&cfg.noop, "noop", false, "don't execute pandoc (show options)")
	flag.BoolVar(&cfg.renumber, "r", false, "")
	flag.BoolVar(&cfg.renumber, "renumber", false, "renumber markdown source files")
	flag.BoolVar(&cfg.help, "?", false, "")
	flag.BoolVar(&cfg.help, "help", false, "displays this help message")
	flag.BoolVar(&cfg.version, "v", false, "")
	flag.BoolVar(&cfg.version, "version", false, "print version and exit")
	return cfg
}

func main() {
	log.SetFlags(0)
	cfg := initFlags()
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: "+name+` [OPTIONS]

Wrapper around pandoc to use a templatable options file.

OPTIONS:
  -c, --check <dir>
          check image directory for orphans
  -n, --noop
          don't execute pandoc (show options)
  -r, --renumber
          renumber markdown source files
  -?, --help
          display this help message
  -v, --version
          print version and exit

EXAMPLES:`)

		fmt.Fprintln(os.Stderr, "\n  $ "+name+` --noop
		
	Example output here...`)
	}
	flag.Parse()

	if flag.Arg(0) == "version" || cfg.version {
		fmt.Printf("%s %s, built on %s (commit: %s)\n", name, version, date, commit)
		return
	}
	if cfg.help {
		flag.Usage()
		return
	}
	if cfg.imagePath != "" {
		listOrphans(cfg.imagePath)
	}
	if cfg.renumber {
		renumberFiles()
	}
	runPandoc(cfg.noop)
}
