package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Functions

func ReadFile(filename string) string {
	// Get the path of the file that called this function, for debugging purposes
	_, callername, _, ok := runtime.Caller(1)
	if !ok {
		panic("Could not find caller")
	}
	dir := filepath.Dir(callername)

	path := filepath.Join(dir, filename)
	data, err := os.ReadFile(path)
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

// 2d grid

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.Y, p.X)
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func GetNextPoint(p Point, d Direction) Point {
	switch d {
	case North:
		p.Y--
	case East:
		p.X++
	case South:
		p.Y++
	case West:
		p.X--
	}
	return p
}

func GetGridValue(grid [][]string, p Point) (string, error) {
	if p.Y < 0 || p.Y >= len(grid) || p.X < 0 || p.X >= len(grid[p.Y]) {
		return "", fmt.Errorf("index out of bounds")
	}
	return grid[p.Y][p.X], nil
}

func GetGridFromString(str string) [][]string {
	lines := strings.Split(str, "\n")
	grid := make([][]string, len(lines))

	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}

	return grid
}

func GetGridPositionByValue(grid [][]string, val string) Point {
	var start Point
	for y, row := range grid {
		for x, p := range row {
			if p == val {
				start = Point{X: x, Y: y}
				break
			}
		}

		if start != (Point{}) {
			break
		}
	}

	// if start == (Point{}) {
	// 	return start, errors.New("Point does not exist")
	// }

	return start
}

func HasDuplicates(slice1, slice2 []int) bool {
	elements := make(map[int]bool)

	for _, val := range slice1 {
		elements[val] = true
	}

	for _, val := range slice2 {
		if elements[val] {
			return true
		}
	}

	return false
}

func Run(part int, function func() int) {
	start := time.Now()
	output := function()
	duration := time.Since(start)

	green := "\033[32m"
	reset := "\033[0m"

	fmt.Printf("Part %d: \t%s%d%s \t(Execution time: %s)\n", part, green, output, reset, duration)
}
