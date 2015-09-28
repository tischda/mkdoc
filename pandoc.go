package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

const dateFormat = "%d.%02d.%02d~-~%02d:%02d:%02d"
const metadataFileName = `metadata.yaml`
const optionsFileName = `pandoc.options`

type metadata struct {
	Target string `yaml:"target"`
	Tag    string
	Date   string
	Time   string
}

// read options and run pandoc on 99-filename.md files
func runPandoc() {
	pandocOptions := readOptionsFile(optionsFileName)
	options := strings.Fields(fillMeta(pandocOptions))

	inputFiles := getMarkdownFileList()
	inputFiles = append(inputFiles, metadataFileName)
	options = append(options, inputFiles...)

	executeProcess("pandoc", options...)
}

// fill pandoc options template with document metadata
// 		Target: obtained from 'metadata.yaml'
// 		Tag: 	'git describe --tags'
// 		Date: 	current date
// 		Time: 	current time
func fillMeta(template string) string {
	meta := readFileMetadata(metadataFileName)

	meta.Tag = getGitTag()

	t := time.Now()
	meta.Date = formatDate(t)
	meta.Time = formatTime(t)

	return fillTemplate(template, meta)
}

// return git tag from project directory
func getGitTag() string {
	cmdOut, _ := exec.Command("git", "describe", "--tags").Output()
	return strings.TrimSpace(string(cmdOut))
}

// return formatted date
func formatDate(t time.Time) string {
	return fmt.Sprintf("%02d.%02d.%d", t.Day(), t.Month(), t.Year())
}

// return formatted time
func formatTime(t time.Time) string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}

// fill text template with data
func fillTemplate(text string, data interface{}) string {
	template, err := template.New("test").Parse(text)
	checkFatal(err)

	buff := bytes.NewBufferString("")
	err = template.Execute(buff, data)
	checkFatal(err)

	return buff.String()
}
