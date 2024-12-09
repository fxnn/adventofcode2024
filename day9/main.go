package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fxnn/adventofcode2024/util"
)

const FREE = -1

func calculateDiskBlocks(diskMap string) []int {
	var diskBlocks = []int{}

	var freeSpace = false
	var fileId = 0
	for _, b := range diskMap {
		var blockCount = util.Atoi(string(b))
		var blockUsage int
		if freeSpace {
			blockUsage = FREE
		} else {
			blockUsage = fileId
			fileId++
		}
		diskBlocks = append(diskBlocks, util.Times(blockCount, blockUsage)...)
		freeSpace = !freeSpace
	}

	return diskBlocks
}

func printDiskBlocks(diskBlocks []int) {
	for _, b := range diskBlocks {
		if b == FREE {
			fmt.Print(".")
		} else {
			fmt.Printf("%d", b)
		}
	}
}

func compact(in []int) []int {
	var out = make([]int, len(in))
	var j = len(out) - 1
	for i := range in {
		for j >= 0 && in[j] == FREE {
			j--
		}
		if i > j {
			out[i] = FREE
		} else if in[i] == FREE {
			out[i] = in[j]
			j--
		} else {
			out[i] = in[i]
		}
	}
	return out
}

func checksum(diskBlocks []int) int {
	var checksum = 0
	for i, b := range diskBlocks {
		if b != FREE {
			checksum += i * b
		}
	}
	return checksum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	var diskMap = scanner.Text()

	var diskBlocks = calculateDiskBlocks(diskMap)
	printDiskBlocks(diskBlocks)
	fmt.Println()

	diskBlocks = compact(diskBlocks)
	printDiskBlocks(diskBlocks)
	fmt.Println()

	var checksum = checksum(diskBlocks)
	fmt.Printf("checksum: %d\n", checksum)
}
