package main

import (
	"math"
	"strings"

	"github.com/banjo/advent-of-code-2024/utils"
)

func calculateIfSafe(levels []int) bool {
	if levels[1] == levels[0] {
		return false
	}

	increasing := levels[1] > levels[0]
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

func part2() int {
	return 0
}

func main() {
	utils.Run(1, part1)
	utils.Run(2, part2)
}
