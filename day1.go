package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func convertAndAppend(list []int, strVal string) []int {
	intVal, err := strconv.Atoi(strVal)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't convert %s to int: %s\n", strVal, err)
		os.Exit(1)
	}
	return append(list, intVal)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func count(list []int, val int) int {
	count := 0
	for _, elem := range list {
		if elem == val {
			count++
		}
	}
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	pattern := regexp.MustCompile(`^(\d+)\s+(\d+)$`)

	var list1 []int
	var list2 []int
	for scanner.Scan() {
		line := scanner.Text()
		matches := pattern.FindStringSubmatch(line)
		if len(matches) > 0 {
			list1 = convertAndAppend(list1, matches[1])
			list2 = convertAndAppend(list2, matches[2])
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	sort.Ints(list1)
	sort.Ints(list2)
	distanceSum := 0
	for i := 0; i < len(list1); i++ {
		distanceSum += abs(list1[i] - list2[i])
	}

	fmt.Printf("sum of distances is %d\n", distanceSum)

	similarity := 0
	for i := 0; i < len(list1); i++ {
		similarity += list1[i] * count(list2, list1[i])
	}
	fmt.Printf("similarity score is %d\n", similarity)
}
