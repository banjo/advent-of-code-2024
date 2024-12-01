package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

func ReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}

	return string(data)
}

type FunctionType func() int

func Run(part int, function FunctionType) {
	start := time.Now()
	output := function()
	duration := time.Since(start)
	green := "\033[32m"
	reset := "\033[0m"

	fmt.Printf("Part %d: \t%s%d%s \t(Execution time: %s)\n", part, green, output, reset, duration)
}
