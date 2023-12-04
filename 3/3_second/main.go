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

	gearregexp, _ := regexp.Compile("\\*")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		gearsIndex := gearregexp.FindAllStringIndex(line, -1)

		for j := 0; j < len(gearsIndex); j++ {
			ratio := findRatio(lines, i-1, gearsIndex[j][0]-1, i+1, gearsIndex[j][1])
			sum = sum + ratio
		}
	}

	fmt.Println("Sum is ", sum)
}

func findRatio(lines []string, ub int, leb int, lb int, rib int) int {
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

	numregexp, _ := regexp.Compile("\\d+")
	partNumbers := 0
	ratio := 1
	for i := ub; i <= lb; i++ {
		nums := numregexp.FindAllString(lines[i], -1)
		numsIndex := numregexp.FindAllStringIndex(lines[i], -1)
		for j := 0; j < len(nums); j++ {
			start := numsIndex[j][0]
			end := numsIndex[j][1] - 1
			if end < leb {
				continue
			}
			if start > rib {
				continue
			}

			num, _ := strconv.Atoi(nums[j])
			ratio = ratio * num
			partNumbers++
		}
	}
	if partNumbers > 1 {
		return ratio
	}
	return 0
}
