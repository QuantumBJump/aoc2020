package main

import (
	"fmt"
	"log"

	"github.com/QuinnStevens/aoc2020/helpers"
)

type set struct {
	m map[rune]struct{}
}

// NewSet creates a new set of runes
func NewSet() *set {
	s := &set{}
	s.m = make(map[rune]struct{})
	return s
}

// Add adds the given rune to the set.
// Returns true if a change was made, false otherwise.
func (s *set) Add(value rune) {
	s.m[value] = exists
}

func (s *set) Remove(value rune) {
	delete(s.m, value)
}

func (s *set) Contains(value rune) bool {
	_, c := s.m[value]
	return c
}

func (s *set) Length() int {
	return len(s.m)
}

func union(set1 *set, set2 *set) *set {
	output := NewSet()

	for item := range set1.m {
		if set2.Contains(item) {
			output.Add(item)
		}
	}

	return output
}

var exists = struct{}{}

func puzzle1(input []string) {
	fmt.Printf("\nPUZZLE 1:\n")
	yesAnswers := NewSet()
	totalYes := 0
	groupYes := 0

	for _, line := range input {
		// For each line in the input
		if line == "" {
			// If the line is blank, we've reached the end of a group, we can start counting the next one.
			totalYes += groupYes  // Add to the total across all groups
			groupYes = 0          // Reset current group's total
			yesAnswers = NewSet() // Delete current group
		} else {
			for _, char := range line {
				// Go through the line one character at a time, & check if it's already been counted.
				new := !yesAnswers.Contains(char)
				yesAnswers.Add(char)
				if new {
					groupYes++
				}
			}
		}
	}

	// Fix OBOE
	totalYes += groupYes

	fmt.Printf("\tTotal yeses: %d\n", totalYes)
}

// PUZZLE 2
func puzzle2(input []string) {
	fmt.Println()
	fmt.Println("Puzzle 2:")
	fmt.Println()

	unanimousAnswers := NewSet()
	total := 0
	noInGroup := 0 // The index of the current questioner in group.
	for _, line := range input {
		if line == "" {
			// Group is ended
			total += unanimousAnswers.Length()
			noInGroup = 0 // Reset group counter
			unanimousAnswers = NewSet()
		} else {
			if noInGroup == 0 {
				// If first in group, populate possible unanimous answers.
				for _, char := range line {
					unanimousAnswers.Add(char)
				}
			} else {
				personAnswers := NewSet()
				for _, char := range line {
					personAnswers.Add(char)
				}
				unanimousAnswers = union(unanimousAnswers, personAnswers)
			}
			noInGroup++
		}
	}
	total += unanimousAnswers.Length()
	fmt.Printf("\nTotal answers: %d\n", total)
}

func main() {
	fmt.Println("Day 6:")
	fmt.Println()

	input, err := helpers.ReadInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	puzzle1(input)
	puzzle2(input)
}
