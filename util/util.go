package util

import "fmt"

// Abs calculates the absolute value of an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Sign returns -1, 0 or 1 depending on the sign of the given integer
func Sign(x int) int {
	if x > 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	return -1
}

// RemoveElement removes element with given index from slice
func RemoveElement(slice []int, index int) ([]int, error) {
	if index < 0 || index >= len(slice) {
		return nil, fmt.Errorf("index out of range")
	}

	// Create a new slice and copy the elements before the index
	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)

	// Append the elements after the index
	newSlice = append(newSlice, slice[index+1:]...)

	return newSlice, nil
}
