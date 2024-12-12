package main

import (
	"fmt"
	"strings"

	"github.com/banjo/advent-of-code-2024/utils"
)

func createBlocks(input string) []*int {
	chars := strings.Split(input, "")
	var res []*int
	for idx, entry := range chars {
		isFile := idx%2 == 0
		fileId := idx / 2
		val := utils.ToInt(entry)

		var toAppend *int
		if isFile {
			toAppend = &fileId
		} else {
			toAppend = nil
		}

		for range val {
			res = append(res, toAppend)
		}

	}

	return res
}

func fragment(s []*int) []int {
	end := len(s) - 1
	start := 0

	for {
		if end < start {
			break
		}

		startVal := s[start]

		if startVal != nil {
			start++
			continue
		}

		for {
			if end < start {
				break
			}

			endVal := s[end]

			if endVal == nil {
				end--
				continue
			}

			s[start] = endVal
			s[end] = nil
			start++
			end--
			break
		}

	}

	a := utils.PointerArrayToIntArray(s)
	return a
}

func getChecksum(s []int) int {
	res := 0
	for idx, id := range s {
		blockValue := id * idx
		res += blockValue
	}

	return res
}

func part1() int {
	content := utils.ReadFile("./input.txt")
	blocks := createBlocks(content)
	updatedBlocks := fragment(blocks)
	checksum := getChecksum(updatedBlocks)
	return checksum
}

func print(blocks []*int) {
	for _, val := range blocks {
		if val == nil {
			fmt.Print("nil")
		} else {
			fmt.Print(*val)
		}
		fmt.Print(" ")
	}

	fmt.Println("")
}

func part2() int {
	// content := utils.ReadFile("./input.txt")
	// fmt.Println("part 2")
	return 0
}

func main() {
	utils.Run(1, part1)
	utils.Run(2, part2)
}
