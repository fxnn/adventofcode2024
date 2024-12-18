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
			var val string
			switch space[y][x] {
			case CORRUPTED:
				val = "C"
			case UNIVSITED:
				val = "."
			default:
				val = util.Itoa(space[y][x])
			}
			if max > 9 {
				fmt.Printf("%3s ", val)
			} else {
				fmt.Printf("%2s", val)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func readPoints() []util.Point {
	var scanner = bufio.NewScanner(os.Stdin)
	var pattern = regexp.MustCompile(`^(\d+),(\d+)$`)

	var points = []util.Point{}
	for scanner.Scan() {
		var line = scanner.Text()
		var matches = pattern.FindStringSubmatch(line)
		var x, y int
		if len(matches) > 0 {
			x = util.Atoi(matches[1])
			y = util.Atoi(matches[2])
			points = append(points, util.Point{Y: y, X: x})
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	return points
}

func resetSpace(space [][]int) {
	for y := range space {
		for x := range space[y] {
			if space[y][x] > 0 {
				space[y][x] = 0
			}
		}
	}
}

func fillSpace(space [][]int, points []util.Point) {
	for _, point := range points {
		space[point.Y][point.X] = CORRUPTED
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

	var points = readPoints()
	var space = makeSpace(*max)
	fillSpace(space, points[:*count])
	printSpace(space)
	var steps = searchPath(space) - 1
	fmt.Printf("found exit in %d steps\n", steps)
	printSpace(space)

	for _, point := range points[*count:] {
		resetSpace(space)
		fillSpace(space, []util.Point{point})
		if searchPath(space) == 0 {
			fmt.Printf("no way out at coordinate %d,%d\n", point.X, point.Y)
			break
		}
	}
}
