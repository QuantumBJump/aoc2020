package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/QuinnStevens/aoc2020/helpers"
)

func puzzle1(input []int) int {
	defer helpers.TimeTracker(time.Now(), "Puzzle 1")
	fmt.Println("Puzzle 1:")
	for i := 25; i < len(input); i++ {
		preamble := input[i-25 : i]
		number := input[i]
		valid := false
		for j, num1 := range preamble {
			if valid {
				break
			}
			for _, num2 := range preamble[j+1:] {
				if num1+num2 == number {
					valid = true
					break
				}
			}
		}
		if !valid {
			// We found an invalid one, we have our answer.
			fmt.Printf("  Index %v (%v) is invalid!\n", i, number)
			return number
		}
	}
	return -1
}

func sum(input []int) int {
	total := 0
	for _, item := range input {
		total += item
	}
	return total
}

func largest(input []int) int {
	largest := input[0]
	for _, item := range input {
		if item > largest {
			largest = item
		}
	}
	return largest
}

func smallest(input []int) int {
	smallest := input[0]
	for _, item := range input {
		if item < smallest {
			smallest = item
		}
	}
	return smallest
}
func puzzle2(input []int, key int) {
	defer helpers.TimeTracker(time.Now(), "Puzzle 2")
	fmt.Println("Puzzle 2:")
	for i := range input {
		// go through each number at a time
		for j := 1; j < len(input[i+1:]); j++ {
			subslice := input[i : i+j]
			if sum(subslice) == key {
				// We've found the subslice!
				result := largest(subslice) + smallest(subslice)
				fmt.Printf("  Found weakness: %v\n", result)
				return
			}
		}
	}
}
func main() {
	strinput, err := helpers.ReadInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	input := []int{}
	for _, item := range strinput {
		num, err := strconv.Atoi(item)
		if err != nil {
			log.Fatalf("Error parsing number %s: %e", item, err)
		}
		input = append(input, num)
	}

	key := puzzle1(input)

	puzzle2(input, key)
}
