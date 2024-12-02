package main

import (
	"math"
	"strings"

	"github.com/banjo/advent-of-code-2024/utils"
)

func calculateIfSafe(levels []int) bool {
	first := levels[0]
	second := levels[1]
	third := levels[2]

	if first == second {
		return false
	}

	increasing := first < second
	nextIncreasing := second < third

	if increasing != nextIncreasing {
		return false
	}

	for idx, level := range levels[1:] {
		previousLevel := levels[idx]
		diff := int(math.Abs(float64(level - previousLevel)))

		if diff < 1 || diff > 3 {
			return false
		}

		nextIsIncreasing := level > previousLevel
		if increasing != nextIsIncreasing {
			return false
		}

		increasing = nextIsIncreasing
	}

	return true
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func part1() int {
	content := utils.ReadFile("./input.txt")
	lines := strings.Split(content, "\n")

	safeCount := 0
	for _, line := range lines {
		strLevels := strings.Fields(line)
		levels := utils.MapStringArrayToIntArray(strLevels)
		isSafe := calculateIfSafe(levels)

		if isSafe {
			safeCount++
		}
	}

	return safeCount
}

func createPossibleCombos(slice []int) [][]int {
	res := make([][]int, len(slice))
	for idx := range slice {
		res[idx] = removeIndex(slice, idx)
	}
	return res
}

func isComboSafe(combos [][]int) bool {
	for _, combo := range combos {
		if calculateIfSafe(combo) {
			return true
		}
	}
	return false
}

func part2() int {
	content := utils.ReadFile("./input.txt")
	lines := strings.Split(content, "\n")

	safeCount := 0
	safes := make([][]int, 0)
	for _, line := range lines {
		strLevels := strings.Fields(line)
		levels := utils.MapStringArrayToIntArray(strLevels)
		possibleCombos := createPossibleCombos(levels)
		isSafe := isComboSafe(possibleCombos)

		if isSafe {
			safes = append(safes, levels)
			safeCount++
		}

	}

	return safeCount
}

func main() {
	utils.Run(1, part1)
	utils.Run(2, part2)
}
