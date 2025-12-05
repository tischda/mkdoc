package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// get list of files in directory
func getFileListInDir(dirName string) []string {
	dir, err := os.Open(dirName)
	if err != nil {
		log.Fatalln(err)
	}
	defer dir.Close() // nolint:errcheck
	fileNames, err := dir.Readdirnames(-1)
	if err != nil {
		log.Fatalln(err)
	}
	sort.Strings(fileNames)
	return fileNames
}

// read metadata from pandoc yaml header file
func readFileMetadata(fileName string) *metadata {
	meta := &metadata{}
	if err := yaml.Unmarshal(fs.readFile(fileName), meta); err != nil {
		log.Fatalln(err)
	}
	return meta
}

// return file contents as string, discard commented lines
func readOptionsFile(name string) string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close() // nolint:errcheck

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cl := scanner.Text()
		if !strings.HasPrefix(cl, "#") {
			lines = append(lines, cl)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return strings.Join(lines, "\n")
}

// read all files into a single string
func mergeFilesToBuffer() string {
	inputFiles := fs.getMarkdownFileList()
	var buffer []byte
	for _, path := range inputFiles {
		data := fs.readFile(path)
		buffer = append(buffer, data...)
	}
	return string(buffer)
}
