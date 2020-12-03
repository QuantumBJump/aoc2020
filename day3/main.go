package main

import (
	"fmt"
	"log"

	"github.com/QuinnStevens/aoc2020/helpers"
)

func parseInput(input []string) [][]string {
	hill := [][]string{}
	for _, line := range input {
		filledLine := []string{}
		for _, character := range line {
			filledLine = append(filledLine, string(character))
		}
		hill = append(hill, filledLine)
	}
	return hill
}

func path(input [][]string, slopeX int, slopeY int) int {
	hill := make([][]string, len(input))
	for i := range hill {
		hill[i] = make([]string, len(input[i]))
	}
	println(copy(hill, input))
	x := 0
	y := 0
	collisions := 0
	for {
		x += slopeX
		y += slopeY
		// We've reached the bottom!
		if y >= len(hill) {
			break
		}
		// Wrap around from the right to the left
		if x >= len(hill[y]) {
			x = x - len(hill[y])
		}
		// check if we've hit a tree and mark our position for debugging
		if hill[y][x] == "#" {
			collisions++
		}
	}

	return collisions
}

func printHill(input [][]string) {
	for _, line := range input {
		fmt.Printf("\t")
		for _, char := range line {
			fmt.Printf("%s", char)
		}
		fmt.Printf("\n")
	}
}

func puzzle1(input [][]string, slopeX int, slopeY int) {
	fmt.Printf("\tpuzzle 1:\n")
	collisions := path(input, slopeX, slopeY)

	fmt.Printf("\n\tCollisions: %d\n\n", collisions)

}

func puzzle2(input [][]string) {
	fmt.Printf("\tpuzzle 2:\n")
	results := []int{}
	product := 1
	collisions := path(input, 1, 1)
	fmt.Printf("\t%d\n\n", collisions)
	results = append(results, collisions)

	collisions = path(input, 3, 1)
	fmt.Printf("\t%d\n\n", collisions)
	results = append(results, collisions)

	collisions = path(input, 5, 1)
	fmt.Printf("\t%d\n\n", collisions)
	results = append(results, collisions)

	collisions = path(input, 7, 1)
	fmt.Printf("\t%d\n\n", collisions)
	results = append(results, collisions)

	collisions = path(input, 1, 2)
	fmt.Printf("\t%d\n\n", collisions)
	results = append(results, collisions)

	fmt.Printf("\n")
	for _, item := range results {
		fmt.Printf("%d, ", item)
		product = product * item
	}
	fmt.Printf("\tAnswer: %d\n", product)
}

func main() {
	fmt.Println("Day 3:")
	input, err := helpers.ReadInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	hill := parseInput(input)
	printHill(hill)
	puzzle1(hill, 3, 1)

	puzzle2(hill)
}
