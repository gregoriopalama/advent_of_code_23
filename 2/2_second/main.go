package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	setsregexp, _ := regexp.Compile(": (.*)")

	sum := 0

	input, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		sum = sum + getPower(setsregexp.FindStringSubmatch(line)[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum is ", sum)
}

func getPower(sets string) int {
	colors := make(map[string]int)
	colors["red"] = 1
	colors["green"] = 1
	colors["blue"] = 1
	for _, set := range strings.Split(sets, "; ") {
		for _, cubes := range strings.Split(set, ", ") {
			color, quantity := extractColorAndQuantity(cubes)
			if quantity > colors[color] {
				colors[color] = quantity
			}
		}
	}
	return colors["red"] * colors["green"] * colors["blue"]
}

func extractColorAndQuantity(cubes string) (string, int) {
	informations := strings.Split(cubes, " ")
	quantity, _ := strconv.Atoi(informations[0])
	return informations[1], quantity
}
