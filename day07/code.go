package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/banjo/advent-of-code-2024/utils"
)

type Parsed struct {
	nums []int
	sum  int
}

func parse(content string) []Parsed {
	lines := strings.Split(content, "\n")

	var results []Parsed
	for _, line := range lines {
		parts := strings.Split(line, ":")
		sum := utils.ToInt(parts[0])
		nums := utils.MapStringArrayToIntArray(strings.Fields(parts[1]))

		result := Parsed{sum: sum, nums: nums}
		results = append(results, result)
	}

	return results
}

func generateBinaryPossibilities(n int) []string {
	var possibilities []string
	total := 1 << n

	for i := 0; i < total; i++ {
		binaryStr := fmt.Sprintf("%0*b", n, i)
		possibilities = append(possibilities, binaryStr)
	}

	return possibilities
}

func part1() int {
	content := utils.ReadFile("./input.txt")
	parsed := parse(content)

	totalSum := 0
	for _, p := range parsed {
		sum := p.sum

		operatorPlaces := len(p.nums) - 1
		possibilities := generateBinaryPossibilities(operatorPlaces)

		for _, binary := range possibilities {
			result := p.nums[0]
			for idx, num := range p.nums[1:] {
				if binary[idx] == '0' {
					result += num
				} else {
					result *= num
				}
			}

			if result == sum {
				totalSum += sum
				break
			}
		}

	}

	return totalSum
}

func generateTernaryPossibilities(n int) []string {
	var possibilities []string
	total := int(math.Pow(3, float64(n)))

	for i := 0; i < total; i++ {
		ternaryStr := ""
		num := i
		for j := 0; j < n; j++ {
			ternaryStr = fmt.Sprintf("%d", num%3) + ternaryStr
			num /= 3
		}
		possibilities = append(possibilities, ternaryStr)
	}

	return possibilities
}

func part2() int {
	content := utils.ReadFile("./input.txt")
	parsed := parse(content)

	totalSum := 0
	for _, p := range parsed {
		sum := p.sum

		operatorPlaces := len(p.nums) - 1
		possibilities := generateTernaryPossibilities(operatorPlaces)

		for _, ternary := range possibilities {
			result := p.nums[0]
			for idx, num := range p.nums[1:] {
				switch ternary[idx] {
				case '0':
					result += num
				case '1':
					result *= num
				case '2':
					resStr := utils.ToString(result)
					numStr := utils.ToString(num)
					result = utils.ToInt(resStr + numStr)
				}
			}

			if result == sum {
				totalSum += sum
				break
			}
		}

	}

	return totalSum
}

type IntOperation func(int, int) int

func traverse(target int, sum int, remaining []int, operations []IntOperation) bool {
	if len(remaining) == 0 {
		return sum == target
	}

	if sum > target {
		return false
	}

	upcoming := remaining[0]
	newRemaining := remaining[1:]
	isCorrect := false
	for _, o := range operations {
		newSum := o(sum, upcoming)
		matches := traverse(target, newSum, newRemaining, operations)
		if matches {
			isCorrect = true
		}
	}

	return isCorrect
}

func part1Traverse() int {
	content := utils.ReadFile("./input.txt")
	parsed := parse(content)

	add := func(a, b int) int { return a + b }
	multiply := func(a, b int) int { return a * b }
	operations := []IntOperation{add, multiply}

	sum := 0
	for _, p := range parsed {
		success := traverse(p.sum, 0, p.nums, operations)
		if success {
			sum += p.sum
		}
	}
	return sum
}

func part2Traverse() int {
	content := utils.ReadFile("./input.txt")
	parsed := parse(content)

	add := func(a, b int) int { return a + b }
	multiply := func(a, b int) int { return a * b }
	combine := func(a, b int) int {
		resStr := utils.ToString(a)
		numStr := utils.ToString(b)
		result := utils.ToInt(resStr + numStr)
		return result
	}
	operations := []IntOperation{add, multiply, combine}

	sum := 0
	for _, p := range parsed {
		success := traverse(p.sum, 0, p.nums, operations)
		if success {
			sum += p.sum
		}
	}
	return sum
}

func main() {
	fmt.Println("Brute")
	utils.Run(1, part1)
	utils.Run(2, part2)

	fmt.Println("\nRecursive")
	utils.Run(1, part1Traverse)
	utils.Run(2, part2Traverse)
}
