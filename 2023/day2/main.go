package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file-path>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	input, err := readInput(filePath)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Solve the problem
	resultPart1 := solvePart1(input)
	fmt.Println("Part 1:", resultPart1)

	resultPart2 := solvePart2(input)
	fmt.Println("Part 2:", resultPart2)
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func solvePart1(input []string) int {
	return 0
}

func solvePart2(input []string) int {
	return 0
}
