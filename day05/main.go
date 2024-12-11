package main

import (
	"strings"

	"github.com/banjo/advent-of-code-2024/utils"
)

func getRulesAndUpdates(content string) ([][]int, [][]int, map[int][]int) {
	lines := strings.Split(content, "\n")

	isRule := true

	var rules [][]int
	var updates [][]int
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
			updates = append(updates, p)
		}

	}

	return rules, updates, rulesMap
}

func isValidUpdate(update []int, rulesMap map[int][]int) bool {
	for idx, page := range update {
		shouldGoBefore := rulesMap[page]
		numbersAfter := update[idx+1:]

		if utils.HasDuplicates(shouldGoBefore, numbersAfter) {
			return false
		}
	}

	return true
}

func part1() int {
	content := utils.ReadFile("./input.txt")
	_, updates, rulesMap := getRulesAndUpdates(content)

	count := 0
	for _, pages := range updates {
		isValid := isValidUpdate(pages, rulesMap)

		if isValid {
			middle := pages[len(pages)/2]
			count += middle
		}

	}

	return count
}

func part2() int {
	content := utils.ReadFile("./input.txt")
	_, updates, rulesMap := getRulesAndUpdates(content)

	var notValid [][]int
	for _, update := range updates {
		isValid := isValidUpdate(update, rulesMap)

		if !isValid {
			notValid = append(notValid, update)
		}

	}

	count := 0
	for _, update := range notValid {

		var res []int

		buffer := make([]int, len(update))
		copy(buffer, update)

		for len(buffer) > 0 {
			page := buffer[0]
			numbersAfter := buffer[1:]
			shouldGoBefore := rulesMap[page]

			if !utils.HasDuplicates(shouldGoBefore, numbersAfter) {
				res = append(res, page)
				buffer = buffer[1:]
			} else {
				buffer = buffer[1:]
				buffer = append(buffer, page)
			}

		}

		middle := res[len(res)/2]
		count += middle
	}

	return count
}

func main() {
	utils.Run(1, part1)
	utils.Run(2, part2)
}
