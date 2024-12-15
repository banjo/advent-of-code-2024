package main

import (
	"fmt"

	"github.com/banjo/advent-of-code-2024/utils"
)

func filterPointsByRules(current utils.Point, possible []utils.Point) []utils.Point {
	var res []utils.Point
	expectedValue := 1

	if current.Value != nil {
		currentValue := utils.ToInt(*current.Value)
		expectedValue = currentValue + 1
	}

	for _, p := range possible {
		if p.Value == nil {
			continue
		}

		if *p.Value == "." {
			continue
		}

		if utils.ToInt(*p.Value) != expectedValue {
			continue
		}

		res = append(res, p)
	}
	return res
}

func part1(file string) int {
	content := utils.ReadFile(file)
	grid := utils.GetGridFromString(content)
	trailheads := utils.GetGridPositionsByValue(grid, "0")

	score := 0
	for _, p := range trailheads {
		scoreVisits := make(map[string]int)

		var buffer []utils.Point
		buffer = append(buffer, p)

		for len(buffer) > 0 {
			current := buffer[0]
			val := utils.ToInt(*current.Value)

			if val == 9 {
				scoreVisits[current.String()] += 1
				buffer = buffer[1:]
				continue
			}

			allPoints := utils.GetPossibleNextPoints(current)
			validPoints := utils.FilterValidPointsInGrid(grid, allPoints)

			filtered := filterPointsByRules(current, validPoints)
			buffer = append(buffer, filtered...)

			buffer = buffer[1:]
		}

		score += len(scoreVisits)

	}

	return score
}

func part2(file string) int {
	// content := utils.ReadFile(file)
	fmt.Println(file)
	return 0
}

func main() {
	utils.Run(1, func() int {
		return part1("./input.txt")
	})
	// utils.Run(2, func() int {
	// 	return part2("./input.txt")
	// })
}
