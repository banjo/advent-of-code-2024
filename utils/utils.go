package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(data))
}

func WriteFile(filename string, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func ToInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
}

func MapStringArrayToIntArray(strs []string) []int {
	ints := make([]int, len(strs))

	for i, level := range strs {
		ints[i] = ToInt(level)
	}

	return ints
}

func Run(part int, function func() int) {
	start := time.Now()
	output := function()
	duration := time.Since(start)

	green := "\033[32m"
	reset := "\033[0m"

	fmt.Printf("Part %d: \t%s%d%s \t(Execution time: %s)\n", part, green, output, reset, duration)
}
