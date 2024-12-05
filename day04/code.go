package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/banjo/advent-of-code-2024/utils"
)

func countMatches(str string, words []string) int {
	matchLength := len(words[0])

	res := 0
	for i := 0; i <= len(str)-matchLength; i++ {
		subString := str[i : i+matchLength]

		for _, word := range words {
			if subString == word {
				res++
			}
		}
	}

	return res
}

func getKey(y int, x int) string {
	return fmt.Sprintf("%d,%d", y, x)
}

func getDiagonalDownRightParentKey(y int, x int) string {
	diff := math.Min(float64(y), float64(x))
	return getKey(y-int(diff), x-int(diff))
}

func getDiagonalUpRightParentKey(y int, x int, colLength int) string {
	diff := math.Min(float64(y), float64(colLength-1-x))
	return getKey(y-int(diff), x+int(diff))
}

func part1() int {
	content := utils.ReadFile("./input.txt")
	lines := strings.Split(content, "\n")
	patterns := []string{"XMAS", "SAMX"}

	verticalStrings := make(map[int]string)
	diagonalDownRightStrings := make(map[string]string)
	diagonalUpRightStrings := make(map[string]string)

	totalMatches := 0
	for yIndex, line := range lines {

		// horizontal strings
		totalMatches += countMatches(line, patterns)

		chars := strings.Split(line, "")
		for xIndex, char := range chars {
			parentKey := getKey(yIndex, xIndex)
			verticalStrings[xIndex] += char

			if xIndex == 0 {
				// left line for down right
				diagonalDownRightStrings[parentKey] += char
			} else if yIndex == 0 {
				// top line for both
				diagonalDownRightStrings[parentKey] += char
			} else {

				// all other ones
				downRightKey := getDiagonalDownRightParentKey(yIndex, xIndex)
				diagonalDownRightStrings[downRightKey] += char
			}
		}

		// for up right diagonal
		for xIndex, char := range chars {
			parentKey := getKey(yIndex, xIndex)

			if xIndex == len(chars)-1 {
				// right line for down right
				diagonalUpRightStrings[parentKey] += char
			} else if yIndex == 0 {
				// top line
				diagonalUpRightStrings[parentKey] += char
			} else {
				// all other ones
				upRightKey := getDiagonalUpRightParentKey(yIndex, xIndex, len(chars))
				diagonalUpRightStrings[upRightKey] += char
			}
		}
	}

	for _, verticalString := range verticalStrings {
		totalMatches += countMatches(verticalString, patterns)
	}

	for _, str := range diagonalDownRightStrings {
		totalMatches += countMatches(str, patterns)
	}

	for _, str := range diagonalUpRightStrings {
		totalMatches += countMatches(str, patterns)
	}

	return totalMatches
}

func isWord(str string) bool {
	return str == "SAM" || str == "MAS"
}

func part2() int {
	content := utils.ReadFile("./input.txt")
	grid := utils.GetGridFromString(content)
	maxY := len(grid) - 1
	maxX := len(grid[0]) - 1

	isOut := func(val int, dir string) bool {
		if val < 0 {
			return true
		}

		if dir == "y" && val > maxY {
			return true
		} else if dir == "x" && val > maxX {
			return true
		}

		return false
	}

	isMatch := func(y, x int) bool {
		if isOut(y+1, "y") || isOut(y-1, "y") || isOut(x+1, "x") || isOut(x-1, "x") {
			return false
		}

		first := grid[y-1][x-1] + grid[y][x] + grid[y+1][x+1]
		second := grid[y+1][x-1] + grid[y][x] + grid[y-1][x+1]

		return isWord(first) && isWord(second)
	}

	count := 0
	for y, row := range grid {
		for x := range row {
			success := isMatch(y, x)

			if success {
				count++
			}
		}
	}

	return count
}

func main() {
	utils.Run(1, part1)
	utils.Run(2, part2)
}
