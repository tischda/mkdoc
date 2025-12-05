package main

import (
	"log"
	"regexp"
)

// listOrphans checks for images in the given path that are not referenced
// in any of the markdown files. If it finds any, it prints them to the log
// and terminates the program.
func listOrphans(path string) {
	orphans, ok := allImagesUsed(path)
	if !ok {
		log.Println("Orphan images:")
		for _, name := range orphans {
			log.Println("  ", name)
		}
		log.Fatal("Fatal: Image check failed.")
	}
}

// allImagesUsed checks if all images in the specified directory path are
// referenced in the markdown files. It returns a slice of unreferenced
// image names (orphans) and a boolean that is true if all images are used,
// and false otherwise.
func allImagesUsed(path string) ([]string, bool) {
	ok := true
	orphans := []string{}
	haystack := mergeFilesToBuffer()
	images := getFileListInDir(path)
	for _, name := range images {
		exp := regexp.MustCompile(name)
		if !exp.MatchString(haystack) {
			orphans = append(orphans, name)
			ok = false
		}
	}
	return orphans, ok
}
