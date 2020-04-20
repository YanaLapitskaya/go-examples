package main

import "fmt"

// Implement average([6]int) float64 function that returns an average value of array (sum / N)
func average(arr [6]int) (result float64) {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return float64(sum) / float64(len(arr))
}

func printAverage(arr [6]int) {
	averageValue := average(arr)
	fmt.Printf("For array %v, average value will be %v.\n", arr, averageValue)
}

func main() {
	// array from task description, expected value - 3.5
	arr := [6]int{1, 2, 3, 4, 5, 6}
	printAverage(arr)

	// array of random values, expected value - 4
	arr = [6]int{3, 85, 3, 22, 9, -98}
	printAverage(arr)
}
