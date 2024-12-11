package main

import (
	"fmt"
	"strings"

	"github.com/banjo/advent-of-code-2024/utils"
)

type Antenna struct {
	val string
	y   int
	x   int
}

func (p Antenna) String() string {
	return fmt.Sprintf("(%s,%d,%d)", p.val, p.y, p.x)
}

type AntennaMap map[string][]Antenna

func parse(content string) AntennaMap {
	antennaMap := make(map[string][]Antenna)
	for y, line := range strings.Split(content, "\n") {
		row := strings.Split(line, "")
		for x, char := range row {
			if char != "." {
				a := Antenna{y: y, x: x, val: char}
				antennaMap[char] = append(antennaMap[char], a)
			}
		}
	}

	return antennaMap
}

func calculateAntinodes(a, b Antenna) []utils.Point {
	yDiff := a.y - b.y
	xDiff := a.x - b.x
	aPoint := utils.Point{Y: b.y - yDiff, X: b.x - xDiff}
	bPoint := utils.Point{Y: a.y + yDiff, X: a.x + xDiff}
	return []utils.Point{aPoint, bPoint}
}

func calculateManyAntinodes(a, b Antenna, amount int) []utils.Point {
	yDiff := a.y - b.y
	xDiff := a.x - b.x
	firstStart := utils.Point{X: b.x, Y: b.y}
	antinodes := []utils.Point{firstStart}
	for range amount {
		latest := antinodes[len(antinodes)-1]
		point := utils.Point{Y: latest.Y - yDiff, X: latest.X - xDiff}
		antinodes = append(antinodes, point)
	}

	secondStart := utils.Point{X: a.x, Y: a.y}
	antinodes = append(antinodes, secondStart)
	for range amount {
		latest := antinodes[len(antinodes)-1]
		point := utils.Point{Y: latest.Y + yDiff, X: latest.X + xDiff}
		antinodes = append(antinodes, point)
	}

	return antinodes
}

func part1() int {
	content := utils.ReadFile("./input.txt")
	grid := utils.GetGridFromString(content)
	antennaMap := parse(content)

	allAntinodes := make(map[utils.Point]bool)
	for _, antennas := range antennaMap {
		buffer := make([]Antenna, len(antennas))
		copy(buffer, antennas)

		for len(buffer) > 0 {
			current := buffer[0]
			rest := buffer[1:]

			for _, sibling := range rest {
				antinodes := calculateAntinodes(current, sibling)
				for _, a := range antinodes {
					_, err := utils.GetGridValue(grid, a)
					if err == nil {
						allAntinodes[a] = true
					}
				}
			}

			buffer = rest
		}
	}

	return len(allAntinodes)
}

func part2() int {
	content := utils.ReadFile("./input.txt")
	grid := utils.GetGridFromString(content)
	antennaMap := parse(content)

	allAntinodes := make(map[utils.Point]bool)
	for _, antennas := range antennaMap {
		buffer := make([]Antenna, len(antennas))
		copy(buffer, antennas)

		for len(buffer) > 0 {
			current := buffer[0]
			rest := buffer[1:]

			for _, sibling := range rest {
				antinodes := calculateManyAntinodes(current, sibling, 50)
				for _, a := range antinodes {
					_, err := utils.GetGridValue(grid, a)
					if err == nil {
						allAntinodes[a] = true
					}
				}
			}

			buffer = rest
		}
	}

	return len(allAntinodes)
}

func main() {
	utils.Run(1, part1)
	utils.Run(2, part2)
}

