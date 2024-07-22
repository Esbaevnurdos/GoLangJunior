package main

import "fmt"

func main() {
	// numbers := []int{1, 10, 100}
	sum := sumup(1, 10, 100)
	fmt.Println(sum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}
