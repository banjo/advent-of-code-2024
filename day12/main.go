package main

import (
	"github.com/banjo/advent-of-code-2024/utils"
)

type Garden struct {
	entries   []utils.Point
	area      int
	perimeter int
}

func part1(file string) int {
	content := utils.ReadFile(file)
	grid := utils.GetGridFromString(content)
	visited := make(map[string]int)

	pathFind := func(p utils.Point) Garden {
		points := []utils.Point{p}

		garden := Garden{entries: points, area: 0, perimeter: 0}

		var buffer []utils.Point
		buffer = append(buffer, p)
		expected := *p.Value

		for len(buffer) > 0 {
			current := buffer[0]
			buffer = buffer[1:]

			visited[current.String()]++

			if *current.Value != expected {
				continue
			}

			connecting := utils.GetPointsAroundWithValue(grid, current)
			pointsOnBoard := utils.FilterValidPointsInGrid(grid, connecting)

			var upcomingConnecting []utils.Point
			for _, p := range pointsOnBoard {

				if *p.Value != expected {
					continue
				}

				if visited[p.String()] > 0 {
					continue
				}

				if utils.ContainsByStringEq(buffer, p) {
					continue
				}

				upcomingConnecting = append(upcomingConnecting, p)
			}

			perimeter := 0
			for _, p := range connecting {
				if p.Value == nil || *p.Value != expected {
					perimeter++
				}
			}

			garden.perimeter += perimeter
			garden.area++
			points = append(points, current)

			buffer = append(buffer, upcomingConnecting...)
		}

		garden.entries = points
		return garden
	}

	cost := 0
	handler := func(p utils.Point) {
		if visited[p.String()] > 0 {
			return
		}

		garden := pathFind(p)
		cost += garden.area * garden.perimeter
	}

	utils.IterateGrid(grid, handler)

	return cost
}

func part2(file string) int {
	// _ := utils.ReadFile(file)
	// fmt.Println("part 2")
	return 0
}

func main() {
	utils.Run(1, func() int {
		return part1("./input.txt")
	})
	utils.Run(2, func() int {
		return part2("./input.txt")
	})
}
