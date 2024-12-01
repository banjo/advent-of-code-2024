package utils

import (
	"fmt"
	"os"
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

func Run(part int, function func() int) {
	start := time.Now()
	output := function()
	duration := time.Since(start)

	green := "\033[32m"
	reset := "\033[0m"

	fmt.Printf("Part %d: \t%s%d%s \t(Execution time: %s)\n", part, green, output, reset, duration)
}
