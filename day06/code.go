package main

import (
	"github.com/banjo/advent-of-code-2024/utils"
)

func nextDirection(current utils.Direction) utils.Direction {
	return (current + 1) % 4
}

func part1() int {
	content := utils.ReadFile("./input.txt")
	grid := utils.GetGridFromString(content)
	start := utils.GetGridPositionByValue(grid, "^")

	visited := make(map[string]bool)
	buffer := []utils.Point{start}
	direction := utils.North

	for len(buffer) > 0 {
		current := buffer[0]
		visited[current.String()] = true

		next := utils.GetNextPoint(current, direction)
		val, err := utils.GetGridValue(grid, next)
		if err != nil {
			break
		}

		if val == "#" {
			direction = nextDirection(direction)
			next = utils.GetNextPoint(current, direction)
		}
		buffer = append(buffer[1:], next)
	}

	return len(visited)
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
