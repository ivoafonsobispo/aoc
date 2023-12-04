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

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
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

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	fmt.Println(sum)
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
