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

// get sorted list of Markdown input files from current directory
func getMarkdownFileList() []string {
	fileList, _ := filepath.Glob("[0-9][0-9]*.md")
	sort.Strings(fileList)
	return fileList
}

// get list of files in directory
func getFileListInDir(dirName string) []string {
	dir, err := os.Open(dirName)
	checkFatal(err)
	fileNames, err := dir.Readdirnames(-1)
	checkFatal(err)
	return fileNames
}

// read metadata from file (we only need 'Target')
func readFileMetadata(fileName string) *metadata {
	meta := &metadata{}
	yaml.Unmarshal(readFile(fileName), meta)
	return meta
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

// read all files into a single string
func mergeFilesToBuffer() string {
	inputFiles := getMarkdownFileList()
	var buffer []byte
	for _, path := range inputFiles {
		data := readFile(path)
		buffer = append(buffer, data...)
	}
	return string(buffer)
}

// return file contents as array of bytes
func readFile(name string) []byte {
	data, err := ioutil.ReadFile(name)
	checkFatal(err)
	return data
}

// rename file
func renameFile(from, to string) {
	if from != to {
		fmt.Printf("  %s --> %s\n", from, to)
		err := os.Rename(from, to)
		checkFatal(err)
	}
}
