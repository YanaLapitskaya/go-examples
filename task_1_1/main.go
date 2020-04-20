package main

import (
	"fmt"
	"math"
	"sort"
)

func GetMedian(i []int) float64 {
	sort.Ints(i)
	var medianIndex float64 = float64(len(i)-1) / 2
	aIndex := int(math.Floor(medianIndex))
	bIndex := int(math.Ceil(medianIndex))
	return float64(i[aIndex]+i[bIndex]) / 2
}

func printMedian(arr []int) {
	result := GetMedian(arr)
	fmt.Printf("For array %v, median is %v \n", arr, result)
}

func main() {
	// even test example, expected result - 3.5
	arr := []int{1, 2, 3, 4, 5, 6}
	printMedian(arr)

	// odd test example, expected result - 3
	arr = []int{1, 2, 3, 4, 5}
	printMedian(arr)

	// not sorted test example, expected result still 3
	arr = []int{3, 2, 1, 4, 5}
	printMedian(arr)
}
