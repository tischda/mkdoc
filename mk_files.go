package main

import (
	"bufio"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
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
	fileList := []string{}
	filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		matched, err := filepath.Match("[0-9][0-9]*.md", f.Name())
		if matched {
			fileList = append(fileList, path)
		}
		return err
	})
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
