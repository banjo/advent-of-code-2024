package main

import (
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

			l := len(utils.ToString(stone))
			isEven := l%2 == 0

			if isEven {
				half := l / 2
				firstHalf := utils.ToString(stone)[:half]
				secondHalf := utils.ToString(stone)[half:]

				updated = append(updated, utils.ToInt(firstHalf))
				updated = append(updated, utils.ToInt(secondHalf))
				continue
			}

			updated = append(updated, stone*2024)
		}

		s = updated
	}

	return len(s)
}

func part2(_ string) int {
	// content := utils.ReadFile(file)
	// fmt.Println("part 2")
	return 0
}

func main() {
	utils.Run(1, func() int {
		content := utils.ReadFile("./input.txt")
		return part1(content)
	})
	utils.Run(2, func() int {
		content := utils.ReadFile("./example.txt")
		return part2(content)
	})
}
