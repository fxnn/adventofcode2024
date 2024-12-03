package main

import (
	"bufio"
	"fmt"
	"github.com/fxnn/adventofcode2024/util"
	"os"
	"strconv"
	"strings"
)

func levels(report string) []int {
	strVals := strings.Split(report, " ")
	levels := make([]int, len(strVals))
	for i := range strVals {
		var err error
		levels[i], err = strconv.Atoi(strVals[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error in Atoi: %s\n", err)
			os.Exit(1)
		}
	}
	return levels
}

func isUnsafe(levels []int, index int) bool {
	firstDiff := levels[0] - levels[1]
	diff := levels[index-1] - levels[index]
	if diff == 0 || util.Abs(diff) > 3 {
		fmt.Printf("  unsafe diff (%d) at index %d\n", diff, index)
		return true
	}
	if util.Sign(diff) != util.Sign(firstDiff) {
		fmt.Printf("  unsafe sign (%d,%d) at index %d\n", util.Sign(diff), util.Sign(firstDiff), index)
		return true
	}
	return false
}

func firstUnsafeIndex(levels []int) int {
	for i := range levels {
		if i > 0 {
			if isUnsafe(levels, i) {
				return i
			}
		}
	}
	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	safeReports := 0
	dampenedSafeReports := 0
	for scanner.Scan() {
		report := scanner.Text()
		levels := levels(report)
		fmt.Printf("report %v\n", report)

		index := firstUnsafeIndex(levels)
		if index == -1 {
			safeReports++
			dampenedSafeReports++
			fmt.Printf("  safe\n")
		} else {
			var err error
			levels, err = util.RemoveElement(levels, index)
			fmt.Printf("  dampened %v\n", levels)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error in RemoveElement: %s\n", err)
				os.Exit(1)
			}
			if firstUnsafeIndex(levels) == -1 {
				dampenedSafeReports++
				fmt.Printf("  safe with dampener\n")
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("safe reports: %d\n", safeReports)
	fmt.Printf("dampened safe reports: %d\n", dampenedSafeReports)
}
