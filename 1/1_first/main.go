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
	firstdigitregexp, _ := regexp.Compile("^([a-z]*)([0-9])")
	lastdigitregexp, _ := regexp.Compile("([0-9])([a-z]*)$")
	sum := 0

	input, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		firstdigit := firstdigitregexp.FindStringSubmatch(line)[2]
		lastdigit := lastdigitregexp.FindStringSubmatch(line)[1]

		f, _ := strconv.Atoi(firstdigit)
		l, _ := strconv.Atoi(lastdigit)
		sum = sum + (f * 10) + l
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum is ", sum)
}
