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
	fmt.Println()
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

func calculateChecksum(diskBlocks []int) int {
	var checksum = 0
	for i, b := range diskBlocks {
		if b != FREE {
			checksum += i * b
		}
	}
	return checksum
}

func findFreeSpace(blocks []int, size int) int {
	var targetBegin, foundSize int
	targetBegin = -1
	for i := range blocks {
		if blocks[i] == FREE && targetBegin == -1 {
			targetBegin = i
		}
		if blocks[i] != FREE && targetBegin != -1 {
			foundSize = i - targetBegin
			if foundSize >= size {
				return targetBegin
			}
			targetBegin = -1
		}
	}
	return -1
}

func set(blocks []int, begin, end, fileId int) {
	for i := begin; i <= end; i++ {
		blocks[i] = fileId
	}
}

func defrag(in []int) []int {
	var out = make([]int, len(in))
	copy(out[:], in[:])

	var fileId, sourceBegin, sourceEnd int
	fileId = -1
	for i := len(in) - 1; i >= 0; i-- {
		if fileId > 0 && fileId != in[i] {
			// case 1: we've exited the file
			var size, targetBegin int

			sourceBegin = i + 1
			size = sourceEnd - sourceBegin + 1

			targetBegin = findFreeSpace(out, size)
			if targetBegin >= 0 && targetBegin < sourceBegin {
				fmt.Printf("  move: %d -> %d (size %d)\n", sourceBegin, targetBegin, size)
				set(out, sourceBegin, sourceEnd, FREE)
				set(out, targetBegin, targetBegin+size-1, fileId)
			} else {
				fmt.Printf("  stay: %d (size %d)\n", sourceBegin, size)
			}
			fileId = -1
		}
		if fileId < 0 {
			if in[i] == FREE {
				// case 2: we're crossing void areas
				continue
			}
			// case 3: we're entering the file
			fileId = in[i]
			sourceEnd = i
		}
	}
	return out
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var diskMap string
	scanner.Scan()
	diskMap = scanner.Text()

	var diskBlocks []int
	diskBlocks = calculateDiskBlocks(diskMap)
	printDiskBlocks(diskBlocks)

	var compactDiskBlocks []int
	var compactChecksum int
	compactDiskBlocks = compact(diskBlocks)
	printDiskBlocks(compactDiskBlocks)
	compactChecksum = calculateChecksum(compactDiskBlocks)
	fmt.Printf("compact checksum: %d\n", compactChecksum)

	var defragDiskBlocks []int
	var defragChecksum int
	defragDiskBlocks = defrag(diskBlocks)
	printDiskBlocks(defragDiskBlocks)
	defragChecksum = calculateChecksum(defragDiskBlocks)
	fmt.Printf("defrag checksum: %d\n", defragChecksum)
}
