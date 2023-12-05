package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numbers = map[string]string{
	"zero":  "z0ero",
	"one":   "o1ne",
	"two":   "t2wo",
	"three": "t3hree",
	"four":  "f4our",
	"five":  "f5ive",
	"six":   "s6ix",
	"seven": "s7even",
	"eight": "e8ight",
	"nine":  "n9ine",
}

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
	sum := 0
	for _, line := range input {
		result, err := strconv.Atoi(strconv.Itoa(getFirstInt(line)) + strconv.Itoa(getLastInt(line)))
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			os.Exit(1)
		}
		sum += result
	}

	return sum
}

func solvePart2(input []string) int {
	sum := 0
	for _, line := range input {
		for k, v := range numbers {
			line = strings.Replace(line, k, v, -1)
		}

		result, err := strconv.Atoi(strconv.Itoa(getFirstInt(line)) + strconv.Itoa(getLastInt(line)))
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			os.Exit(1)
		}
		sum += result
	}

	return sum
}

func getFirstInt(s string) int {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return int(c - '0')
		}
	}
	return -1
}

func getLastInt(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if c := s[i]; c >= '0' && c <= '9' {
			return int(c - '0')
		}
	}
	return -1
}
