package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestCompileNewNames(t *testing.T) {

	// standard case
	testCase(t, map[string]string{
		"04-aa": "01-aa",
		"07-bb": "02-bb",
		"08-cc": "03-cc"})

	// upward remaning 01-bb would overwrite 02-bb
	testCase(t, map[string]string{
		"01-aa": "01-aa",
		"01-bb": "02-bb",
		"02-bb": "03-bb"})

	// downward renaming 03-aa would overwrite 02-aa
	testCase(t, map[string]string{
		"02-aa": "01-aa",
		"03-aa": "02-aa",
		"03-bb": "03-bb"})

	// both up-downward renaming
	testCase(t, map[string]string{
		"00-aa": "01-aa",
		"01-aa": "02-aa",
		"02-aa": "03-aa",
		"05-bb": "04-bb",
		"06-bb": "05-bb"})
}

func testCase(t *testing.T, expected map[string]string) {
	inputFiles := getSortedKeys(expected)
	actual, twoStep := compileNewNames(inputFiles)
	if twoStep {
		renamedFiles := mockRename(inputFiles, actual)
		actual, twoStep = compileNewNames(renamedFiles)
		if twoStep {
			t.Errorf("Should not happen:", actual)
		}
	}
	if !reflect.DeepEqual(getSortedValues(actual), getSortedValues(expected)) {
		t.Errorf("Expected: %q, was: %q", expected, actual)
	}
}

// return list of sorted keys from map
func getSortedKeys(mymap map[string]string) []string {
	keys := make([]string, 0, len(mymap))
	for k, _ := range mymap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// return list of sorted values from map
func getSortedValues(mymap map[string]string) []string {
	values := make([]string, 0, len(mymap))
	for _, v := range mymap {
		values = append(values, v)
	}
	sort.Strings(values)
	return values
}

// simulate file renaming
func mockRename(inputFiles []string, newNames map[string]string) []string {
	outputFiles := make([]string, 0, len(inputFiles))
	for _, name := range inputFiles {
		newName, exists := newNames[name]
		if exists {
			// fmt.Printf("%s --> %s\n", name, newName)
			outputFiles = append(outputFiles, newName)
		} else {
			outputFiles = append(outputFiles, name)
		}
	}
	return outputFiles
}
