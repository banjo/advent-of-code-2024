package main

import (
	"fmt"
	"strings"

	"github.com/banjo/advent-of-code-2024/utils"
)

func part1(content string) int {
	strs := strings.Fields(content)
	s := utils.MapStringArrayToIntArray(strs)

	for i := 0; i < 25; i++ {
		var updated []int
		for _, stone := range s {
			if stone == 0 {
				updated = append(updated, 1)
				continue
			}

			if hasEvenDigits(stone) {
				splitted := split(stone)
				updated = append(updated, splitted...)
				continue
			}

			updated = append(updated, stone*2024)
		}

		s = updated
	}

	return len(s)
}

func hasEvenDigits(i int) bool {
	l := len(utils.ToString(i))
	isEven := l%2 == 0
	return isEven
}

func split(i int) []int {
	s := utils.ToString(i)
	l := len(s)
	half := l / 2
	firstHalf := utils.ToString(i)[:half]
	secondHalf := utils.ToString(i)[half:]
	return []int{utils.ToInt(firstHalf), utils.ToInt(secondHalf)}
}

func getKey(stone int, blinks int) string {
	return fmt.Sprintf("%d-%d", stone, blinks)
}

func blink(stone int, blinks int, memo map[string]int) int {
	if blinks == 0 {
		return 1
	}

	key := getKey(stone, blinks)

	if prev, exists := memo[key]; exists {
		return prev
	}

	if stone == 0 {
		res := blink(1, blinks-1, memo)
		memo[key] = res
		return res
	}

	if hasEvenDigits(stone) {
		splitted := split(stone)

		res1 := blink(splitted[0], blinks-1, memo)
		res2 := blink(splitted[1], blinks-1, memo)

		t := res1 + res2
		memo[key] = t
		return t
	}

	res := blink(stone*2024, blinks-1, memo)
	memo[key] = res
	return res
}

func part2(content string) int {
	strs := strings.Fields(content)
	s := utils.MapStringArrayToIntArray(strs)
	memo := make(map[string]int)

	count := 0
	for _, stone := range s {
		count += blink(stone, 75, memo)
	}

	return count
}

func main() {
	utils.Run(1, func() int {
		content := utils.ReadFile("./input.txt")
		return part1(content)
	})
	utils.Run(2, func() int {
		content := utils.ReadFile("./input.txt")
		return part2(content)
	})
}
