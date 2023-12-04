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
	matches := []int{1}

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		m := getmatches(line)

		for j := 0; j < m; j++ {
			if len(matches) <= (i + j + 1) {
				matches = append(matches, 1)
			}
			index := i + j + 1
			matches[index] = matches[index] + matches[i]
		}
		i++
		if len(matches) <= i {
			matches = append(matches, 1)
		}
	}

	for k := 0; k < i; k++ {
		sum = sum + matches[k]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum is ", sum)
}

func getmatches(line string) int {
	winning := strings.Split(strings.Trim(strings.Split(
		strings.Split(line, ":")[1], "|")[0], " "), " ")
	card := strings.Split(strings.Trim(strings.Split(
		strings.Split(line, ":")[1], "|")[1], " "), " ")

	matches := 0

	for _, n := range card {
		if n == "" {
			continue
		}
		if slices.Contains(winning, n) {
			matches++
		}
	}

	return matches
}
