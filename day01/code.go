package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"

	"github.com/banjo/advent-of-code-2024/utils"
)

func part1() int {
	content := utils.ReadFile("input.txt")
	lines := strings.Split(content, "\n")

	var first, second []int

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if trimmedLine == "" {
			continue
		}

		var int1, int2 int
		_, err := fmt.Sscanf(line, "%d %d", &int1, &int2)
		if err != nil {
			log.Println("Error: ", err)
			os.Exit(1)
		}

		first = append(first, int1)
		second = append(second, int2)
	}

	slices.Sort(first)
	slices.Sort(second)

	val := 0
	for i := 0; i < len(first); i++ {
		val1 := first[i]
		val2 := second[i]

		diff := int(math.Abs(float64(val1) - float64(val2)))
		val += diff
	}

	return val
}

func part2() int {
	content := utils.ReadFile("input.txt")
	lines := strings.Split(content, "\n")

	var first, second []int

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if trimmedLine == "" {
			continue
		}

		var int1, int2 int
		_, err := fmt.Sscanf(line, "%d %d", &int1, &int2)
		if err != nil {
			log.Println("Error: ", err)
			os.Exit(1)
		}

		first = append(first, int1)
		second = append(second, int2)
	}

	hashmap := make(map[int]int)
	for _, num := range first {
		if _, exists := hashmap[num]; !exists {
			hashmap[num] = 0
		}
	}

	for _, val := range second {
		if _, exists := hashmap[val]; exists {
			hashmap[val]++
		}
	}

	finalVal := 0
	for _, val := range first {
		occurence := hashmap[val]
		finalVal += occurence * val
	}

	return finalVal
}

func main() {
	utils.Run(1, part1)
	utils.Run(2, part2)
}
