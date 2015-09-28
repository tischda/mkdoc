package main

import (
	"fmt"
	"log"
	"strings"
)

const SUFFIX = ".$$.md"
const MAX_FILES = 99

// Reorder files on disk. Expects a list of files with format "[0-9][0-9]*.md".
// Renaming may require a two step process if renaming would overwrite files.
// This case is handled by adding a suffix and renaming twice.
func renumberFiles() {
	inputFiles := getMarkdownFileList()
	if len(inputFiles) > MAX_FILES {
		message := fmt.Sprintf("Too many files: %d (max %d)", len(inputFiles), MAX_FILES)
		log.Fatal(message)
	}
	fmt.Println("Renumbering files")
	newNames, twoStep := compileNewNames(inputFiles)
	for from, to := range newNames {
		renameFile(from, to)
	}
	if twoStep {
		// now names are unique, remove suffix and ignore twoStep
		newNames, _ = compileNewNames(getMarkdownFileList())
		for from, to := range newNames {
			renameFile(from, to)
		}
	}
	fmt.Println("")
}

// Renumber file names sequentially. Returns a mapping between source file names
// and destination file names. Source and destination values may be the same.
// If the new name of a renumberd file exists in the source list, a suffix will
// be added and the return value of twoStep set to true. This means that you
// should proceed with a two step rename:
//   step 1: 00-aa --> 01.aa.$$
// then call this function again, and rename:
//   step 2: 01.aa.$$ --> 01.aa
func compileNewNames(inputFiles []string) (mapping map[string]string, twoStep bool) {
	mapping = make(map[string]string, len(inputFiles))
	for i, name := range inputFiles {
		newName := fmt.Sprintf("%02d-%s", i+1, strings.TrimSuffix(name[3:], SUFFIX))
		if newName != name {
			if stringInSlice(newName, inputFiles) {
				// would overwrite existing file, add suffix
				newName = fmt.Sprintf("%s%s", newName, SUFFIX)
				twoStep = true
			}
		}
		mapping[name] = newName
	}
	return
}

// Returns true if string found in slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
