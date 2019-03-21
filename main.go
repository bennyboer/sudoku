package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Display array
	fmt.Printf("All: %v\n", arr)

	// Remove 5
	removedFive := arr[:]
	removedFive[4] = removedFive[len(removedFive) - 1]
	removedFive = removedFive[:len(removedFive) - 1]
	fmt.Printf("Removed 5: %v\n", removedFive)

	// Remove all
	removedAll := arr[:0]
	fmt.Printf("Removed all: %v\n", removedAll)
}
