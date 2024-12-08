package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fxnn/adventofcode2024/util"
)

const (
	ADD = '+'
	MUL = '*'
	CON = '|'
)

var allOperators = []byte{ADD, MUL, CON}

func calculate(operands []int, operators []byte) int {
	var result = operands[0]
	for i := 0; i < len(operators); i++ {
		if operators[i] == ADD {
			result += operands[i+1]
		} else if operators[i] == MUL {
			result *= operands[i+1]
		} else {
			result = util.Atoi(util.Itoa(result) + util.Itoa(operands[i+1]))
		}
	}
	return result
}

func formula(operands []int, operators []byte) string {
	var s = util.Itoa(operands[0])
	for i := 0; i < len(operators); i++ {
		s += fmt.Sprintf(" %s %d", string(operators[i]), operands[i+1])
	}

	return s
}

func modifyOperators(operators []byte, index int, newOperator byte) []byte {
	// we need to copy, because slices are, although passed by value, only shallow copies
	// referring to the same backing array
	var result = make([]byte, len(operators))
	copy(result[:], operators[:])
	result[index] = newOperator
	return result
}

func isOperatorsExistStep(testResult int, operands []int, operators []byte, fixedOperators int) bool {
	if fixedOperators >= len(operators) {
		var match = calculate(operands, operators) == testResult
		if match {
			fmt.Printf("%d: %s\n", testResult, formula(operands, operators))
		}
		return match
	}

	for _, o := range allOperators {
		var newOps = modifyOperators(operators, fixedOperators, o)
		if isOperatorsExistStep(testResult, operands, newOps, fixedOperators+1) {
			return true
		}
	}
	return false
}

func makeOperators(count int) []byte {
	var operators = make([]byte, count)
	for i := 0; i < count; i++ {
		operators[i] = ADD
	}
	return operators
}

func isOperatorsExist(testResult int, operands []int) bool {
	var operators = makeOperators(len(operands) - 1)
	var fixedOperators = 0
	return isOperatorsExistStep(testResult, operands, operators, fixedOperators)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var total = 0
	for scanner.Scan() {
		var line = scanner.Text()
		var parts = strings.SplitN(line, ": ", 2)
		var testResult = util.Atoi(parts[0])
		var operands = util.AtoiList(strings.Split(parts[1], " "))
		if isOperatorsExist(testResult, operands) {
			total += testResult
		}
	}

	fmt.Printf("total calibration result: %d\n", total)
}
