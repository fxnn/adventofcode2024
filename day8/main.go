package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/fxnn/adventofcode2024/util"
)

const ANTENNA_PATTERN = "[a-zA-Z0-9]"

func discoverAntennaPoints(lines [][]byte) map[byte][]util.Point {
	var antennaPattern = regexp.MustCompile(ANTENNA_PATTERN)
	var antennaPoints = make(map[byte][]util.Point)
	for y, line := range lines {
		for x, b := range line {
			if antennaPattern.Match([]byte{b}) {
				var list []util.Point
				if l, ok := antennaPoints[b]; ok {
					list = l
				} else {
					list = make([]util.Point, 0)
				}
				antennaPoints[b] = append(list, util.Point{Y: y, X: x})
			}
		}
	}
	return antennaPoints
}

func discoverAntinodes(antennaPoints map[byte][]util.Point, mapHeight int, mapWidth int) map[util.Point]util.Void {
	var antinodes = make(map[util.Point]util.Void)
	for _, points := range antennaPoints {
		for i, point := range points {
			for j, otherPoint := range points {
				if i == j {
					continue
				}
				antinodes[point] = util.Void{}
				antinodes[otherPoint] = util.Void{}
				var delta = point.Subtract(otherPoint)

				var antinode1 = point.Add(delta)
				for antinode1.IsInBounds(mapHeight, mapWidth) {
					antinodes[antinode1] = util.Void{}
					antinode1 = antinode1.Add(delta)
				}

				var antinode2 = otherPoint.Subtract(delta)
				for antinode2.IsInBounds(mapHeight, mapWidth) {
					antinodes[antinode2] = util.Void{}
					antinode2 = antinode2.Subtract(delta)
				}
			}
		}
	}
	return antinodes
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var lines [][]byte
	for scanner.Scan() {
		lines = append(lines, scanner.Bytes())
	}

	var mapHeight = len(lines)
	var mapWidth = len(lines[0])
	var antennaPoints = discoverAntennaPoints(lines)
	var antinodes = discoverAntinodes(antennaPoints, mapHeight, mapWidth)

	fmt.Printf("unique antinodes: %d\n", len(antinodes))
}
