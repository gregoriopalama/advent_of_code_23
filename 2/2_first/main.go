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
	gameregexp, _ := regexp.Compile("^Game ([0-9]*):")
	setsregexp, _ := regexp.Compile(": (.*)")

	sum := 0
	requirements := make(map[string]int)
	requirements["red"] = 12
	requirements["green"] = 13
	requirements["blue"] = 14

	input, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		game, _ := strconv.Atoi(gameregexp.FindStringSubmatch(line)[1])

		ispossible := isGamePossible(requirements, setsregexp.FindStringSubmatch(line)[1])
		if ispossible {
			sum = sum + game
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum is ", sum)
}

func isGamePossible(requirements map[string]int, sets string) bool {
	for _, set := range strings.Split(sets, "; ") {
		for _, cubes := range strings.Split(set, ", ") {
			color, quantity := extractColorAndQuantity(cubes)
			if quantity > requirements[color] {
				return false
			}
		}
	}
	return true
}

func extractColorAndQuantity(cubes string) (string, int) {
	informations := strings.Split(cubes, " ")
	quantity, _ := strconv.Atoi(informations[0])
	return informations[1], quantity
}
