package util

import (
	"fmt"
	"os"
	"strconv"
)

type Void struct{}

// Itoa converts integer to string
func Itoa(i int) string {
	return strconv.Itoa(i)
}

// Atoi converts string to integer, and exist immediately upon error
func Atoi(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		fmt.Fprintf(os.Stderr, "Error in Atoi(%s): %s\n", s, err)
		os.Exit(1)
		return 0
	} else {
		return i
	}
}

// AtoiList converts a list of strings into a list of integers,
// whereas the result list has the same size as the given one.
func AtoiList(l []string) []int {
	var result = make([]int, len(l))
	for i, s := range l {
		result[i] = Atoi(s)
	}
	return result
}

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
