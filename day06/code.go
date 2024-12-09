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
	content := utils.ReadFile("./input.txt")
	grid := utils.GetGridFromString(content)
	start := utils.GetGridPositionByValue(grid, "^")

	visited := make(map[string]int)
	buffer := []utils.Point{start}
	direction := utils.North
	stepsFromStartForObstruction := 0 // start from 0 to just calculate the length at first run
	currentSteps := 0
	currentObstructionPoint := utils.Point{X: -1, Y: -1}

	reset := func() {
		visited = make(map[string]int)
		buffer = []utils.Point{start}
		direction = utils.North
		stepsFromStartForObstruction++
		currentSteps = 0
		currentObstructionPoint = utils.Point{X: -1, Y: -1}
	}

	obstructionCount := 0
	isDone := false
	firstRun := true
	totalLength := 0
	for !isDone {
		skipVisitCheck := false
		for len(buffer) > 0 {
			current := buffer[0]

			if !skipVisitCheck {
				// stop if second time visiting a second run
				if visited[current.String()] == 4 {
					obstructionCount++
					reset()
					break
				}

				visited[current.String()] += 1
			}

			currentSteps++
			next := utils.GetNextPoint(current, direction)
			val, err := utils.GetGridValue(grid, next)
			if err != nil {

				if firstRun {
					totalLength = currentSteps
					firstRun = false
				} else if totalLength == stepsFromStartForObstruction {
					isDone = true
				}

				reset()
				break
			}

			atObstructionPoint := next == currentObstructionPoint
			if atObstructionPoint && !firstRun {
				direction = nextDirection(direction)
				skipVisitCheck = true
				continue
			}

			notVisited := visited[next.String()] == 0
			if currentSteps == stepsFromStartForObstruction && start != next && notVisited && !firstRun {
				currentObstructionPoint = next
				direction = nextDirection(direction)
				skipVisitCheck = true
				continue
			}

			if val == "#" {
				direction = nextDirection(direction)
				skipVisitCheck = true
				continue
			}

			skipVisitCheck = false
			buffer = append(buffer[1:], next)
		}
	}

	return obstructionCount
}

func main() {
	utils.Run(1, part1)
	utils.Run(2, part2)
}
