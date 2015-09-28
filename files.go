package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"fmt"

	"gopkg.in/yaml.v2"
)

// read metadata from file (we only need 'Target')
func readMeta(fileName string) *metadata {
	meta := &metadata{}
	yaml.Unmarshal(readFile(fileName), meta)
	return meta
}

// return file contents as array of bytes
func readFile(name string) []byte {
	data, err := ioutil.ReadFile(name)
	checkFatal(err)
	return data
}

// get sorted list of Markdown input files from current directory
func getMarkdownInputFiles() []string {
	fileList, _ := filepath.Glob("[0-9][0-9]*.md")
	sort.Strings(fileList)
	return fileList
}

// return file contents as string, but discard lines starting with '#'
func readOptionsFile(name string) string {
	file, err := os.Open(name)
	checkFatal(err)
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cl := scanner.Text()
		if !strings.HasPrefix(cl, "#") {
			lines = append(lines, cl)
		}
	}
	checkFatal(scanner.Err())
	return strings.Join(lines, "\n")
}

// rename file
func renameFile(from, to string) {
	if from != to {
		fmt.Printf("%s --> %s\n", from, to)
		err := os.Rename(from, to)
		checkFatal(err)
	}
}
