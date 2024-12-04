package main

import (
	"bufio"
	"fmt"
	"os"
)

const NEEDLE = "MAS"

type point struct {
	y int
	x int
}

func charAt(lines []string, p point) (byte, bool) {
	if p.y < 0 || p.x < 0 {
		return 0, false
	}
	if p.y >= len(lines) {
		return 0, false
	}
	if p.x >= len(lines[p.y]) {
		return 0, false
	}
	return lines[p.y][p.x], true
}

func generateMatchPoints(at point, dy, dx int) []point {
	result := []point{}
	for i := range NEEDLE {
		p := point{y: at.y + dy*i, x: at.x + dx*i}
		result = append(result, p)
	}
	return result
}

func transpose(p point, dy, dx int) point {
	return point{y: p.y + dy, x: p.x + dx}
}

func matchPointGenerators() []func(point) []point {
	return []func(point) []point{
		func(p point) []point { return generateMatchPoints(transpose(p, -1, -1), 1, 1) },
		func(p point) []point { return generateMatchPoints(transpose(p, 1, 1), -1, -1) },
		func(p point) []point { return generateMatchPoints(transpose(p, -1, 1), 1, -1) },
		func(p point) []point { return generateMatchPoints(transpose(p, 1, -1), -1, 1) },
	}
}

func isAMatch(points []point, lines []string) bool {
	for i, p := range points {
		if c, ok := charAt(lines, p); !ok || c != NEEDLE[i] {
			return false
		}
	}
	return true
}

func countMatchesFrom(source point, lines []string, matches [][]byte) int {
	matchCount := 0
	for _, generator := range matchPointGenerators() {
		matchPoints := generator(source)
		if !isAMatch(matchPoints, lines) {
			continue
		}
		matchCount++
	}
	if matchCount != 2 {
		return 0
	}
	for _, generator := range matchPointGenerators() {
		matchPoints := generator(source)
		if !isAMatch(matchPoints, lines) {
			continue
		}
		for i, p := range matchPoints {
			matches[p.y][p.x] = NEEDLE[i]
		}
	}
	return 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	matches := make([][]byte, len(lines))
	for y := range lines {
		matches[y] = make([]byte, len(lines[y]))
		for x := range lines[y] {
			matches[y][x] = byte('.')
		}
	}

	matchCount := 0
	for y := range lines {
		line := lines[y]
		for x := range line {
			matchCount += countMatchesFrom(point{y: y, x: x}, lines, matches)
		}
	}

	for _, lineMatches := range matches {
		fmt.Printf("%s\n", string(lineMatches))
	}
	fmt.Fprintf(os.Stderr, "matches: %d\n", matchCount)
}
