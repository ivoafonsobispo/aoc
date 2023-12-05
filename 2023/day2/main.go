package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
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
		s := strings.Split(line, ":")

		gameID, err := strconv.Atoi(strings.Split(s[0], " ")[1])
		fmt.Println(gameID)
		if err != nil {
			fmt.Println("Error converting gameID:", err)
			return 0
		}
		rounds := strings.Split(s[1], ";")

		maxRed := 0
		maxGreen := 0
		maxBlue := 0

		for _, round := range rounds {
			cube := strings.Split(round, ",")
			for _, c := range cube {
				str := strings.Split(c, " ")
				value, err := strconv.Atoi(str[1])
				if err != nil {
					fmt.Println("Error converting value:", err)
					return 0
				}
				color := str[2]

				switch color {
				case "red":
					if value > maxRed {
						maxRed = value
					}
				case "green":
					if value > maxGreen {
						maxGreen = value
					}
				case "blue":
					if value > maxBlue {
						maxBlue = value
					}
				}
			}
		}

		sum += (maxRed * maxGreen * maxBlue)
	}
	return sum
}

func solvePart2(input []string) int {
	return 0
}
