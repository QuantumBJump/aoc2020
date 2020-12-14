package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/QuinnStevens/aoc2020/helpers"
)

type policy struct {
	min    int
	max    int
	letter byte
}

type dbEntry struct {
	policy   policy
	password string
}

func puzzle1(input []string) {
	re, err := regexp.Compile("(?P<min>\\d+)-(?P<max>\\d+)\\s(?P<letter>\\w): (?P<password>\\w+)")
	if err != nil {
		log.Fatalf("Error compiling regex: %v", err)
	}

	var valid int
	for _, item := range input {
		match := re.FindStringSubmatch(item)
		result := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
		password := []byte(result["password"])
		min, err := strconv.Atoi(result["min"])
		if err != nil {
			log.Fatalf("Error parsing min: %v", err)
		}
		max, err := strconv.Atoi(result["max"])
		if err != nil {
			log.Fatalf("Error parsing max: %v", err)
		}

		var found int
		for _, char := range password {
			letter := result["letter"][0]
			if char == letter {
				found++
			}
		}
		if found >= min && found <= max {
			valid++
		}
	}
	fmt.Printf("Number of valid passwords: %d\n", valid)
}

func puzzle2(input []string) {
	fmt.Printf("\nPuzzle 2:\n\n")
	re, err := regexp.Compile("(?P<first>\\d+)-(?P<second>\\d+)\\s(?P<letter>\\w): (?P<password>\\w+)")
	if err != nil {
		log.Fatalf("Error compiling regex: %v", err)
	}
	names := re.SubexpNames()
	var valid int

	for _, line := range input {
		match := re.FindStringSubmatch(line)
		result := map[string]string{}
		for i, n := range match {
			if i == 0 {
				result["full"] = n
			} else {
				result[names[i]] = n
			}
		}

		pos1, err := strconv.Atoi(result["first"])
		if err != nil {
			log.Fatalf("Error parsing pos1: %v", err)
		}
		pos1--

		pos2, err := strconv.Atoi(result["second"])
		if err != nil {
			log.Fatalf("Error parsing pos2: %v", err)
		}
		pos2--

		letter := byte(result["letter"][0])

		matches := 0
		char1 := result["password"][pos1]
		if char1 == letter {
			matches++
		}
		char2 := result["password"][pos2]
		if char2 == letter {
			matches++
		}

		if matches == 1 {
			valid++
		}

	}
	fmt.Printf("\tNumber of valid passwords: %d\n", valid)
}
func main() {
	input, err := helpers.ReadInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	puzzle1(input)
	puzzle2(input)
}
