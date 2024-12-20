package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/fxnn/adventofcode2024/util"
)

const (
	WALL  = '#'
	EMPTY = '.'
	START = 'S'
	END   = 'E'
)

func readCourse() [][]byte {
	var scanner = bufio.NewScanner(os.Stdin)

	var y int
	var course [][]byte
	for scanner.Scan() {
		var line = scanner.Text()
		var row = make([]byte, len(line))
		for x, r := range line {
			row[x] = byte(r)
		}
		course = append(course, row)
		y++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	return course
}

var directions = []util.Point{{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: -1}}

func findMoves(course [][]byte) []util.Point {
	var f = findStart(course)
	var moves = []util.Point{f}
	for course[f.Y][f.X] != END {
		for _, d := range directions {
			var t = f.Add(d)
			if len(moves) > 1 && t == moves[len(moves)-2] {
				continue
			}
			if !t.IsInBounds(len(course), len(course[f.Y])) {
				continue
			}
			if course[t.Y][t.X] == WALL {
				continue
			}
			moves = append(moves, t)
			//fmt.Printf("move %d,%d -> %d,%d\n", f.Y, f.X, t.Y, t.X)
			f = t
			break
		}
	}
	return moves
}

func findStart(course [][]byte) util.Point {
	for y := 0; y < len(course); y++ {
		for x := 0; x < len(course[y]); x++ {
			if course[y][x] == START {
				return util.Point{Y: y, X: x}
			}
		}
	}
	fmt.Println("Error: start not found")
	os.Exit(1)
	return util.Point{}
}

func findMoveIdx(moves []util.Point, move util.Point) int {
	for i := range moves {
		if moves[i] == move {
			return i
		}
	}
	return -1
}

type cheat struct {
	start util.Point
	end   util.Point
}

func findCheats(course [][]byte, moves []util.Point, startIdx int, minSaving int, cheatedWalls map[util.Point]util.Void) []cheat {
	var s = moves[startIdx]
	var cheats []cheat
	for _, d := range directions {
		var w = s.Add(d)
		if !w.IsInBounds(len(course), len(course[w.Y])) {
			continue
		}
		if course[w.Y][w.X] != WALL {
			continue
		}
		var e = w.Add(d)
		var ei = findMoveIdx(moves, e)
		// HINT: save 100 picoseconds, but this includes two additional steps
		if ei-startIdx < minSaving+2 {
			// TODO: might there be multiple cheats forward?
			continue
		}
		cheatedWalls[w] = util.Void{}
		fmt.Printf("%d picoseconds saved by cheat: %d,%d -> %d,%d\n", ei-startIdx-2, s.Y, s.X, e.Y, e.X)
		cheats = append(cheats, cheat{start: s, end: e})
	}
	return cheats
}

func printCourseWithCheats(course [][]byte, cheats map[util.Point]util.Void) {
	for y := range course {
		for x := range course[y] {
			if _, ok := cheats[util.Point{Y: y, X: x}]; ok {
				fmt.Print("C")
			} else {
				fmt.Print(string(course[y][x]))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	var minSaving = flag.Int("minSaving", 100, "consider cheats that save at least minSaving picoseconds")
	flag.Parse()

	var course = readCourse()
	var moves = findMoves(course)
	var time = len(moves) - 1
	fmt.Printf("course takes %d picoseconds\n", time)
	var cheats = make(map[cheat]util.Void)
	var cheatedWalls = make(map[util.Point]util.Void)
	for i := range moves {
		for _, c := range findCheats(course, moves, i, *minSaving, cheatedWalls) {
			cheats[c] = util.Void{}
		}
	}
	printCourseWithCheats(course, cheatedWalls)
	fmt.Printf("found %d unique cheats saving >= %d picoseconds", len(cheats), *minSaving)
}
