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

func ToString(n int) string {
	return strconv.Itoa(n)
}

func MapStringArrayToIntArray(strs []string) []int {
	ints := make([]int, len(strs))

	for i, level := range strs {
		ints[i] = ToInt(level)
	}

	return ints
}

func PointerArrayToIntArray(s []*int) []int {
	a := make([]int, len(s))
	for idx, val := range s {
		if val != nil {
			a[idx] = *val
		}
	}
	return a
}

// 2d grid

type Point struct {
	Value *string
	X     int
	Y     int
}

func (p Point) String() string {
	if p.Value == nil {
		return fmt.Sprintf(`(y:%d, x:%d, val:nil)`, p.Y, p.X)
	}
	return fmt.Sprintf(`(y:%d, x:%d, val:%s)`, p.Y, p.X, *p.Value)
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
	p.Value = nil
	return p
}

func GetPointsAround(p Point) []Point {
	var possiblePoints []Point
	for d := North; d <= West; d++ {
		n := GetNextPoint(p, d)
		possiblePoints = append(possiblePoints, n)
	}

	return possiblePoints
}

func GetPointsAroundWithValue(grid [][]string, p Point) []Point {
	var possiblePoints []Point
	for d := North; d <= West; d++ {
		n := GetNextPoint(p, d)
		val, err := GetGridValue(grid, n)
		if err == nil {
			n.Value = &val
		}
		possiblePoints = append(possiblePoints, n)
	}

	return possiblePoints
}

func GetGridValue(grid [][]string, p Point) (string, error) {
	if p.Y < 0 || p.Y >= len(grid) || p.X < 0 || p.X >= len(grid[p.Y]) {
		return "", fmt.Errorf("index out of bounds")
	}
	return grid[p.Y][p.X], nil
}

func FilterValidPointsInGrid(grid [][]string, points []Point) []Point {
	var possiblePoints []Point
	for _, p := range points {
		val, err := GetGridValue(grid, p)
		if err != nil {
			continue
		}

		// set value as well
		p.Value = &val
		possiblePoints = append(possiblePoints, p)
	}

	return possiblePoints
}

func SetPointValueInGrid(grid [][]string, p *Point) {
	val := grid[p.Y][p.X]
	p.Value = &val
}

func GetGridFromString(str string) [][]string {
	lines := strings.Split(str, "\n")
	grid := make([][]string, len(lines))

	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}

	return grid
}

type GridCallback func(p Point)

func IterateGrid(grid [][]string, callback GridCallback) {
	for y, row := range grid {
		for x, val := range row {
			p := Point{Y: y, X: x, Value: &val}
			callback(p)
		}
	}
}

func GetGridPositionByValue(grid [][]string, val string) Point {
	var start Point
	for y, row := range grid {
		for x, p := range row {
			if p == val {
				start = Point{X: x, Y: y, Value: &val}
				break
			}
		}

		if start != (Point{}) {
			break
		}
	}

	return start
}

func GetGridPositionsByValue(grid [][]string, val string) []Point {
	var res []Point
	for y, row := range grid {
		for x, p := range row {
			if p == val {
				res = append(res, Point{X: x, Y: y, Value: &val})
				continue
			}
		}
	}

	return res
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

func ContainsByStringEq(slice []Point, p Point) bool {
	pString := p.String()
	for _, v := range slice {
		if v.String() == pString {
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
