package main

import (
	"fmt"
	"go-exercises/week-2/median"
	"go-exercises/week-2/square"
)

func main() {
	// median task
	arr := []int{1, 2, 3, 4, 5, 6}
	result := median.GetMedian(arr)
	fmt.Printf("For array %v, median is %v \n", arr, result)

	// square task
	s := square.New(1, 1, 5)
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
