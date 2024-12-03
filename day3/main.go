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

	pattern := regexp.MustCompile(`(do|don't)\(\)|mul\((\d+),(\d+)\)`)
	result := 0
	do := true
	for _, match := range pattern.FindAllStringSubmatch(text, -1) {
		fmt.Printf("%v\n", match)
		if match[1] == "do" {
			do = true
		} else if match[1] == "don't" {
			do = false
		} else if do {
			n1 := util.Atoi(match[2])
			n2 := util.Atoi(match[3])
			result += n1 * n2
		}
	}

	fmt.Printf("result: %d\n", result)
}
