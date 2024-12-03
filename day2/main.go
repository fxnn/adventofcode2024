package main

import (
	"bufio"
	"flag"
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

func isSafeWithIndexRemoved(levels []int, index int) bool {
	dampenedLevels, err := util.RemoveElement(levels, index)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in RemoveElement: %s\n", err)
		os.Exit(1)
	}
	return firstUnsafeIndex(dampenedLevels) == -1
}

func isSafeWithDampening(levels []int, unsafeIndex int) bool {
	// HINT (0, 1): those we use for determining the gradient
	// HINT (i-1, i): these are the ones we compare for their difference and gradient
	dampeningIndices := []int{0, 1, unsafeIndex - 1, unsafeIndex}
	for _, dampeningIndex := range dampeningIndices {
		if isSafeWithIndexRemoved(levels, dampeningIndex) {
			return true
		}
	}
	return false
}

func main() {
	var bruteforceEnabled bool
	flag.BoolVar(&bruteforceEnabled, "bruteforce", false, "enable bruteforcing the solution")
	flag.Parse()

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
		} else if isSafeWithDampening(levels, index) {
			dampenedSafeReports++
			fmt.Printf("  safe with dampener at index %d\n", index)
		} else if bruteforceEnabled {
			for bruteforceIndex := range levels {
				if isSafeWithIndexRemoved(levels, bruteforceIndex) {
					dampenedSafeReports++
					fmt.Printf("  BRUTEFORCED: safe with dampener at index %d (we tried at %d, %d)\n",
						bruteforceIndex, index-1, index)
					break
				}
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
