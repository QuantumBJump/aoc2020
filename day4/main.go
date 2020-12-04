package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/QuinnStevens/aoc2020/helpers"
)

var requiredFields = []string{
	"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
}

// Entry describes a passport entry. it has the following fields
type Entry struct {
	byr string // birth year
	iyr string // issue year
	eyr string // expiry year
	hgt string // height
	hcl string // hair colour
	ecl string // eye colour
	pid string // passport id
	cid string // country id
}

func splitDatum(input string) (string, string, error) {
	a := strings.Split(input, ":")
	if len(a) != 2 {
		return "", "", fmt.Errorf("There are %v items in this datum, there should be 2", len(a))
	}
	return a[0], a[1], nil
}

func parseInput(input []string) []map[string]string {
	entries := []map[string]string{}
	entry := map[string]string{}
	for _, line := range input {
		// For every line in the file
		if line == "" {
			// It's a newline, we've reached the end of the entry. Push it & start a new one
			entries = append(entries, entry)
			entry = map[string]string{}
		} else {
			// We have items on this line, add them to the entry.
			data := strings.Split(line, " ")
			for _, datum := range data {
				key, value, err := splitDatum(datum)
				if err != nil {
					log.Fatalf("Failed to split datum \"%v\": %v", datum, err)
				}
				entry[key] = value
			}
		}
	}
	entries = append(entries, entry)
	return entries
}
func checkNumber(num string, length, min, max int) bool {
	if len(num) != length {
		return false
	}
	x, err := strconv.Atoi(num)
	if err != nil {
		return false
	}
	if x < min || x > max {
		return false
	}
	return true
}

func checkHeight(entry string) bool {
	re := regexp.MustCompile("^(?P<num>\\d+)(?P<unit>cm|in)$")
	match := re.FindStringSubmatch(entry)
	if match == nil {
		// The regex didn't match
		return false
	}

	// Create a map with the named capture groups
	names := re.SubexpNames()
	result := map[string]string{}
	for i, n := range match {
		if i == 0 {
			result["full"] = n
		} else {
			result[names[i]] = n
		}
	}

	if result["unit"] == "cm" {
		return checkNumber(result["num"], 3, 150, 193)
	} else if result["unit"] == "in" {
		return checkNumber(result["num"], 2, 59, 76)
	} else {
		return false
	}
}

func checkHair(entry string) bool {
	re, err := regexp.Compile("^#[0-9a-f]{6}$")
	if err != nil {
		log.Fatalf("Failed to compile regex.")
	}
	return re.MatchString(entry)
}

func checkEyes(entry string) bool {
	result := false
	acceptableColours := []string{
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth",
	}
	for _, colour := range acceptableColours {
		if entry == colour {
			result = true
		}
	}
	return result
}

func checkPID(entry string) bool {
	re := regexp.MustCompile("^[0-9]{9}$")
	return re.MatchString(entry)
}

func checkValid(entry map[string]string, required []string) (bool, error) {
	valid := true
	// birth year
	var result bool
	result = checkNumber(entry["byr"], 4, 1920, 2002)
	fmt.Printf("\tbyr - %v: %v\n", entry["byr"], result)
	valid = valid && result
	result = checkNumber(entry["iyr"], 4, 2010, 2020)
	fmt.Printf("\tiyr - %v: %v\n", entry["iyr"], result)
	valid = valid && result
	result = checkNumber(entry["eyr"], 4, 2020, 2030)
	fmt.Printf("\teyr - %v: %v\n", entry["eyr"], result)
	valid = valid && result
	result = checkHeight(entry["hgt"])
	fmt.Printf("\thgt - %v: %v\n", entry["hgt"], result)
	valid = valid && result
	result = checkHair(entry["hcl"])
	fmt.Printf("\thcl - %v: %v\n", entry["hcl"], result)
	valid = valid && result
	result = checkEyes(entry["ecl"])
	fmt.Printf("\tecl - %v: %v\n", entry["ecl"], result)
	valid = valid && result
	result = checkPID(entry["pid"])
	fmt.Printf("\tpid - %v: %v\n", entry["pid"], result)
	valid = valid && result

	fmt.Printf("\t  Passport is: %v\n", valid)
	return valid, nil
}

func puzzle1(input []string) {
	entries := parseInput(input)

	valid := 0
	for i, entry := range entries {
		fmt.Printf("Passport %d:\n", i)
		if result, _ := checkValid(entry, requiredFields); result == true {
			valid++
		}
	}

	fmt.Printf("Total valid passports: %d\n", valid)

}

func main() {
	fmt.Printf("Day 4:\n")
	input, err := helpers.ReadInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	puzzle1(input)
}
