package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	sum := 0
	input, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		points := getCardPoints(line)
		sum = sum + points
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum is ", sum)
}

func getCardPoints(line string) int {
	winning := strings.Split(strings.Trim(strings.Split(
		strings.Split(line, ":")[1], "|")[0], " "), " ")
	card := strings.Split(strings.Trim(strings.Split(
		strings.Split(line, ":")[1], "|")[1], " "), " ")

	points := 0

	for _, n := range card {
		if n == "" {
			continue
		}
		if slices.Contains(winning, n) {
			if points == 0 {
				points = 1
			} else {
				points = points * 2
			}
		}
	}

	return points
}
