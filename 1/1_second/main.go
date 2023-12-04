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
		firstdigit := firstdigitregexp.FindStringSubmatch(change_literals_to_digits(line, 1))[2]
		lastdigit := lastdigitregexp.FindStringSubmatch(change_literals_to_digits(line, -1))[1]

		f, _ := strconv.Atoi(firstdigit)
		l, _ := strconv.Atoi(lastdigit)
		sum = sum + (f * 10) + l
		fmt.Println(line, change_literals_to_digits(line, 1), change_literals_to_digits(line, -1), firstdigit, lastdigit)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum is ", sum)
}

func change_literals_to_digits(line string, direction int) string {
	literals_map := make(map[string]int)
	literals_map["one"] = 1
	literals_map["two"] = 2
	literals_map["three"] = 3
	literals_map["four"] = 4
	literals_map["five"] = 5
	literals_map["six"] = 6
	literals_map["seven"] = 7
	literals_map["eight"] = 8
	literals_map["nine"] = 9

	i := len(line)
	if direction < 0 {
		i = -1
	}
	literal := ""

	for key, _ := range literals_map {
		if direction > 0 && strings.Index(line, key) > -1 && strings.Index(line, key) < i {
			i = strings.Index(line, key)
			literal = key
		}
		if direction < 0 && strings.LastIndex(line, key) > -1 && strings.LastIndex(line, key) > i {
			i = strings.LastIndex(line, key)
			literal = key
		}
	}
	if literal != "" {
		line = strings.Replace(line, literal, strconv.Itoa(literals_map[literal]), -1)
	}
	return line
}
