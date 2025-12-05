package main

import "sort"

// return list of sorted keys from map
func getSortedKeys(mymap map[string]string) []string {
	keys := make([]string, 0, len(mymap))
	for k := range mymap {
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
