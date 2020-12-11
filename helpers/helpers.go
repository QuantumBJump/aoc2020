package helpers

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

// ReadInput parses the input file into a slice of strings.
func ReadInput(path string) ([]string, error) {
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

// ReadInputInt takes an input file with one integer per line, and parses it into a slice of ints.
func ReadInputInt(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Error parsing number %v: %v", line, err)
		}
		lines = append(lines, num)
	}

	return lines
}

// ParseNamedRegex takes a regex with named capture groups, and returns a map
// of the capture groups & their matches. Returns nil if not matched.
func ParseNamedRegex(regex string, input string) map[string]string {
	re, err := regexp.Compile(regex)
	if err != nil {
		log.Fatalf("Error compiling regex %s: %v", regex, err)
	}
	names := re.SubexpNames()
	match := re.FindStringSubmatch(input)
	if match == nil {
		return nil
	}

	result := map[string]string{}
	for i, n := range match {
		if i == 0 {
			result["full"] = n
		} else {
			result[names[i]] = n
		}
	}
	return result
}

// TimeTracker can be deferred by a function to time how long it takes.
func TimeTracker(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
