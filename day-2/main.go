package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

func parseInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
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
func main() {
	input, err := parseInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	puzzle1(input)
}
