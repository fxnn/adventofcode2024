package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/fxnn/adventofcode2024/util"
)

const (
	CORRUPTED = -1
)

func printSpace(space [][]int) {
	for y := range space {
		for x := range space[y] {
			fmt.Printf("%2d ", space[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	var size = flag.Int("max", 70, "coordinates have range [0,max]")
	var count = flag.Int("count", 1024, "number of bytes that have fallen")
	flag.Parse()
	var space = make([][]int, *size+1)
	for y := range space {
		space[y] = make([]int, *size+1)
	}

	var scanner = bufio.NewScanner(os.Stdin)
	var pattern = regexp.MustCompile(`^(\d+),(\d+)$`)

	var fallen int
	for scanner.Scan() && fallen < *count {
		var line = scanner.Text()
		var matches = pattern.FindStringSubmatch(line)
		var x, y int
		if len(matches) > 0 {
			x = util.Atoi(matches[1])
			y = util.Atoi(matches[2])
			space[y][x] = CORRUPTED
			fallen++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	printSpace(space)
}
