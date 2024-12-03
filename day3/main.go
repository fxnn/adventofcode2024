package main

import (
	"bufio"
	"fmt"
	"github.com/fxnn/adventofcode2024/util"
	"os"
	"regexp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var text string
	for scanner.Scan() {
		text += scanner.Text()
	}

	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	result := 0
	for _, match := range pattern.FindAllStringSubmatch(text, -1) {
		n1 := util.Atoi(match[1])
		n2 := util.Atoi(match[2])
		result += n1 * n2
	}

	fmt.Printf("result: %d\n", result)
}
