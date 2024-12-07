package main

import (
	"strings"

	"github.com/banjo/advent-of-code-2024/utils"
)

func getRulesAndPages(content string) ([][]int, [][]int, map[int][]int) {
	lines := strings.Split(content, "\n")

	isRule := true

	var rules [][]int
	var pages [][]int
	rulesMap := make(map[int][]int)

	for _, line := range lines {
		if line == "" {
			isRule = false
			continue
		}

		if isRule {
			nums := utils.MapStringArrayToIntArray(strings.Split(line, "|"))
			rulesMap[nums[1]] = append(rulesMap[nums[1]], nums[0])
			rules = append(rules, nums)
		} else {
			p := utils.MapStringArrayToIntArray(strings.Split(line, ","))
			pages = append(pages, p)
		}

	}

	return rules, pages, rulesMap
}

func part1() int {
	content := utils.ReadFile("./input.txt")
	_, listOfPages, rulesMap := getRulesAndPages(content)

	count := 0
	for _, pages := range listOfPages {

		isValid := true
		for idx, page := range pages {
			shouldGoBefore := rulesMap[page]
			numbersAfter := pages[idx+1:]

			for _, after := range numbersAfter {
				for _, before := range shouldGoBefore {
					if before == after {
						isValid = false
						break
					}
				}
				if !isValid {
					break
				}
			}
		}

		if isValid {
			middle := pages[len(pages)/2]
			count += middle
		}

	}

	return count
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

