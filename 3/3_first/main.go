package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	sum := 0
	input, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := []string{}

	lineCount := 0
	lineLength := 0
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		if lineLength < len(line) {
			lineLength = len(line)
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	numregexp, _ := regexp.Compile("\\d+")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		nums := numregexp.FindAllString(line, -1)
		numsIndex := numregexp.FindAllStringIndex(line, -1)

		for j := 0; j < len(nums); j++ {
			if isPartNumber(lines, i-1, numsIndex[j][0]-1, i+1, numsIndex[j][1]) {
				num, _ := strconv.Atoi(nums[j])
				sum = sum + num
			}
		}
	}

	fmt.Println("Sum is ", sum)
}

func isPartNumber(lines []string, ub int, leb int, lb int, rib int) bool {
	if ub < 0 {
		ub = 0
	}
	if leb < 0 {
		leb = 0
	}
	if lb >= len(lines) {
		lb = len(lines) - 1
	}
	if rib >= len(lines[0]) {
		rib = len(lines[0]) - 1
	}

	for i := ub; i <= lb; i++ {
		for j := leb; j <= rib; j++ {
			if isSymbol(lines[i][j]) {
				return true
			}
		}
	}
	return false
}

func isSymbol(character uint8) bool {
	return character != 46 && (character < 48 || character > 57)
}
