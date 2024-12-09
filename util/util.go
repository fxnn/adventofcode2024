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
