package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run init.go <day_number>")
		os.Exit(1)
	}

	dayNumber := os.Args[1]
	dayFolder := fmt.Sprintf("day%02s", dayNumber)

	// run input early if program fails
	input := getInput(dayNumber)

	err := os.Mkdir(dayFolder, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		os.Exit(1)
	}

	createFile(filepath.Join(dayFolder, "code.go"), codeTemplate())
	createFile(filepath.Join(dayFolder, "input.txt"), input)

	fmt.Printf("Initialized day %s\n", dayNumber)
}

func loadEnv(name string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Could not find secret: ", name)
		os.Exit(1)
	}

	secret := os.Getenv(name)

	return secret
}

func createFile(filename, content string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
}

func codeTemplate() string {
	return `package main

import (
    "fmt"

    "github.com/banjo/advent-of-code-2024/utils"
)

func part1() int {
		content := utils.ReadFile("./input.txt")
    fmt.Println("part 1")
    return 0
}

func part2() int {
		content := utils.ReadFile("./input.txt")
    fmt.Println("part 2")
    return 0
}

func main() {
    utils.Run(1, part1)
    utils.Run(2, part2)
}`
}

func getInput(day string) string {
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%s/input", day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: loadEnv("AOC_SESSION"),
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Error fetching input: %s\nResponse body: %s\n", resp.Status, string(body))
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	return string(body)
}
