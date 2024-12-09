package util

import "fmt"

func Times[T any](size int, value T) []T {
	a := make([]T, size)
  for i := range a {
		a[i] = value
	}
	return a
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
