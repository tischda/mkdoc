package main

import (
	"log"
	"regexp"
)

func listOrphans() {
	orphans, ok := allImagesUsed()
	if !ok {
		log.Println("Orphan images:")
		for _, name := range orphans {
			log.Println("  ", name)
		}
		log.Fatal("Fatal: Image check failed.")
	}
}

func allImagesUsed() ([]string, bool) {
	ok := true
	orphans := []string{}
	haystack := mergeFilesToBuffer()
	images := getFileListInDir(imgCheckDir)
	for _, name := range images {
		exp := regexp.MustCompile(name)
		if !exp.MatchString(haystack) {
			orphans = append(orphans, name)
			ok = false
		}
	}
	return orphans, ok
}
