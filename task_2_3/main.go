package main

import (
	"fmt"
	"sort"
)

// Implement printSorted(map[int]string) funtcion that prints map values sorted in order of increasing keys
func printSorted(m map[int]string) []string {
	var keys []int
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	result := make([]string, len(keys))
	for index, value := range keys {
		result[index] = m[value]
	}
	return result
}

func main() {
	// example from task, expected - ["b", "c", "a"]
	m := map[int]string{2: "a", 0: "b", 1: "c"}
	fmt.Println(printSorted(m))

	// example from task, expected = ["bb", "aa", "cc"]
	m = map[int]string{10: "aa", 0: "bb", 500: "cc"}
	fmt.Println(printSorted(m))
}
