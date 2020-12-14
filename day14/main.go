package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/QuinnStevens/aoc2020/helpers"
)

var debug bool = true

func puzzle1(input []string) {
	defer helpers.TimeTracker(time.Now(), "Puzzle 1")
	fmt.Println("Puzzle 1:")

	var onemask int64
	var zeromask int64
	memory := map[int]int64{}
	for _, line := range input {
		// Read through the lines of the input
		words := strings.Split(line, " ")
		if words[0] == "mask" {
			// Create masks
			onemask = 0
			zeromask = 0
			for _, char := range words[2] {
				onemask = onemask << 1
				zeromask = zeromask << 1
				if char == '1' {
					onemask++
				} else if char == '0' {
					zeromask++
				}
			}
			if debug {
				fmt.Printf("One mask: %v\n", strconv.FormatInt(onemask, 2))
				fmt.Printf("Zero mask: %v\n", strconv.FormatInt(zeromask, 2))
			}
		} else {
			result := helpers.ParseNamedRegex("^mem\\[(?P<addr>\\d+)\\] = (?P<val>\\d+)$", line)
			address, err := strconv.Atoi(result["addr"])
			if err != nil {
				log.Fatalf("Error parsing address %v: %v", result["addr"], err)
			}
			num, err := strconv.Atoi(result["val"])
			value := int64(num)
			if err != nil {
				log.Fatalf("Error parsing value %v: %v", result["val"], err)
			}

			if debug {
				fmt.Println()
				fmt.Printf("Value: %v\n", strconv.FormatInt(value, 2))
			}
			// Apply masks
			value = value | onemask
			value = value &^ zeromask
			if debug {
				fmt.Printf("Address: %v\n", address)
				fmt.Printf("Value: %v\n", strconv.FormatInt(value, 2))
			}
			memory[address] = value
		}
	}

	total := int64(0)
	for _, value := range memory {
		total += value
	}
	fmt.Printf("  Answer: %v\n", total)
}

func main() {
	input, err := helpers.ReadInput("input.txt")
	if err != nil {
		log.Fatalf("Failed to parse input: %v", err)
	}

	puzzle1(input)
}
