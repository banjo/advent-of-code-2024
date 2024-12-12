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
		if id == -1 {
			continue
		}
		blockValue := id * idx
		res += blockValue
	}

	return res
}

func print(blocks []*int) string {
	s := ""
	for _, val := range blocks {
		if val == nil {
			fmt.Print("nil")
			s += "nil"
		} else {
			fmt.Print(*val)
			s += utils.ToString(*val)
		}
		fmt.Print(" ")
		s += " "
	}

	fmt.Println("")
	return s
}

func printAoCFormat(blocks []int) {
	for _, val := range blocks {
		if val == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(val)
		}
	}

	fmt.Println("")
}

func fragmentByFiles(s []*int) []int {
	end := len(s) - 1
	start := 0

	for end >= 0 {
		endVal := s[end]

		if endVal == nil {
			end--
			continue
		}

		fileSpace := 0
		fileSlice := s[:end+1] // include current
		for idx := range fileSlice {
			indexOfLast := len(fileSlice) - idx - 1
			last := fileSlice[indexOfLast]

			if last != nil && *last == *endVal {
				fileSpace++
				continue
			}
			break
		}

		start = 0
		for {
			if start > len(s)-1 {
				end -= fileSpace
				break
			}

			startVal := s[start]

			if startVal != nil {
				start++
				continue
			}

			freeSpace := 0
			for _, val := range s[start:] {
				if val != nil {
					break
				}

				freeSpace++
			}

			if start+freeSpace >= end {
				end -= fileSpace
				break
			}

			if fileSpace > freeSpace {
				// look for new file block
				start += fileSpace
				continue
			}

			if start > end {
				end -= fileSpace
				break
			}

			fileBlocksToMove := fileSlice[len(fileSlice)-fileSpace:]
			for idx := range fileBlocksToMove {
				startIdx := start + idx
				endIdx := end - idx
				s[startIdx] = endVal
				s[endIdx] = nil
			}

			end--
			break
		}

	}

	a := make([]int, len(s))
	for idx, val := range s {
		if val != nil {
			a[idx] = *val
		} else {
			a[idx] = -1
		}
	}
	return a
}

func part1() int {
	content := utils.ReadFile("./input.txt")
	blocks := createBlocks(content)
	updatedBlocks := fragment(blocks)
	checksum := getChecksum(updatedBlocks)
	return checksum
}

func part2() int {
	// NOT WORKING
	content := utils.ReadFile("./input.txt")
	blocks := createBlocks(content)
	updatedBlocks := fragmentByFiles(blocks)
	checksum := getChecksum(updatedBlocks)
	return checksum
}

func main() {
	utils.Run(1, part1)
	utils.Run(2, part2)
}
