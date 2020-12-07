package helpers

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

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
