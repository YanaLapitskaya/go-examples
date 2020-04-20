package main

import "fmt"

// Write max([]string) function that returns the longest word from the slice of strings (the first if there are no more than one)
func max(arr []string) string {
	// using variable maxLength, so we don't need to call len() in array each time
	maxLength := len(arr[0])
	maxEl := arr[0]
	for _, value := range arr {
		if maxLength < len(value) {
			maxLength = len(value)
			maxEl = value
		}
	}
	return maxEl
}

// Write reverse([]int64) []int64 function that returns the copy of the original slice in reverse order
func reverse(arr []int64) []int64 {
	arrLen := len(arr)
	newArr := make([]int64, arrLen)
	for i := arrLen - 1; i > -1; i-- {
		newArr[arrLen-(i+1)] = arr[i]
	}
	return newArr
}

func main() {
	// example for max, expected value - "three"
	arr1 := []string{"one", "two", "three"}
	fmt.Println(max(arr1))

	// example for max, expected value - "one"
	arr1 = []string{"one", "two"}
	fmt.Println(max(arr1))

	// example for reverse, expected value - (15, 5, 2, 1)
	arr2 := []int64{1, 2, 5, 15}
	fmt.Println(reverse(arr2))
}
