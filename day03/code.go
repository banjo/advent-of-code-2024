package main

import (
	"regexp"

	"github.com/banjo/advent-of-code-2024/utils"
)

func part1() int {
	content := utils.ReadFile("./input.txt")
	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(content, -1)

	val := 0
	for _, match := range matches {
		val += utils.ToInt(match[1]) * utils.ToInt(match[2])
	}

	return val
}

func part2() int {
	content := utils.ReadFile("./input.txt")
	r, _ := regexp.Compile(`(do\(\))|(don't\(\))|(mul\((\d+),(\d+)\))`)
	matches := r.FindAllStringSubmatch(content, -1)

	val := 0
	enabled := true
	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {
			val += utils.ToInt(match[4]) * utils.ToInt(match[5])
		}
	}

	return val
}

func main() {
	utils.Run(1, part1)
	utils.Run(2, part2)
}

