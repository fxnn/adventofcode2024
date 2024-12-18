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
	UNIVSITED = 0
)

func printSpace(space [][]int) {
	var max = len(space) - 1
	for y := range space {
		for x := range space[y] {
			switch space[y][x] {
			case CORRUPTED:
				fmt.Print(" C")
			case UNIVSITED:
				fmt.Print(" .")
			default:
				fmt.Printf("%2d", space[y][x])
			}
			if max > 9 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func readSpace(space [][]int, count int) {
	var scanner = bufio.NewScanner(os.Stdin)
	var pattern = regexp.MustCompile(`^(\d+),(\d+)$`)

	var fallen int
	for scanner.Scan() && fallen < count {
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
}

func makeSpace(size int) [][]int {
	var space = make([][]int, size+1)
	for y := range space {
		space[y] = make([]int, size+1)
	}
	return space
}

var directions = []util.Point{{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: -1}}

func markAround(from util.Point, space [][]int, set []util.Point) []util.Point {
	var size = len(space)
	var steps = space[from.Y][from.X]
	for _, direction := range directions {
		var to = from.Add(direction)
		if to.IsInBounds(size, size) && space[to.Y][to.X] == 0 {
			set = append(set, to)
			space[to.Y][to.X] = steps + 1
		}
	}
	return set
}

func expandSet(set []util.Point, space [][]int) []util.Point {
	var newset = []util.Point{}
	for _, point := range set {
		newset = markAround(point, space, newset)
	}
	return newset
}

func searchPath(space [][]int) int {
	var max = len(space) - 1
	var set []util.Point = []util.Point{{X: 0, Y: 0}}

	space[0][0] = 1
	for space[max][max] == 0 && len(set) > 0 {
		set = expandSet(set, space)
	}

	return space[max][max]
}

func main() {
	var max = flag.Int("max", 70, "coordinates have range [0,max]")
	var count = flag.Int("count", 1024, "number of bytes that have fallen")
	flag.Parse()

	var space = makeSpace(*max)
	readSpace(space, *count)
	printSpace(space)
	var steps = searchPath(space) - 1
	fmt.Printf("found exit in %d steps\n", steps)
	printSpace(space)
}
