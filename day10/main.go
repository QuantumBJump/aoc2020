package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"time"

	"github.com/QuinnStevens/aoc2020/helpers"
)

func puzzle1(input []int) {
	defer helpers.TimeTracker(time.Now(), "Puzzle 1")
	fmt.Println("Puzzle 1:")
	sort.Ints(input)
	// map [voltage difference]number seen
	currVoltage := 0
	differences := map[int]int{}

	differences[1] = 0
	differences[2] = 0
	differences[3] = 0

	for _, num := range input {
		diff := num - currVoltage
		if diff > 3 {
			log.Fatalf("Difference greater than 3! %v", diff)
		} else {
			differences[diff]++
			currVoltage = num
		}
	}
	differences[3]++ // The difference between the final adapter and the laptop is always 3
	fmt.Printf("  Number of 1-jolt differences: %v\n", differences[1])
	fmt.Printf("  Number of 3-jolt differences: %v\n", differences[3])
	fmt.Printf("  Answer: %v\n", differences[1]*differences[3])

}

// Takes a []int sorted in reverse order, so largest is at front
func perms(input []int) int {
	var result int
	if len(input) <= 2 {
		return 1
	}
	for i := 1; i < 4; i++ {
		if i == len(input) {
			break
		}
		if input[0]-input[i] <= 3 {
			result += perms(input[i:])
		}
	}
	return result
}

func fixedPoint(input []int) bool {
	// Takes a slice of ints. Works out if the first int in the slice is a fixed point for a jump length of 3
	if len(input) == 1 {
		return true
	}
	if input[0]-input[1] == 3 {
		return true
	}
	return false
}

func puzzle2(input []int) {
	defer helpers.TimeTracker(time.Now(), "Puzzle 2")
	fmt.Println("Puzzle 2:")
	sort.Sort(sort.Reverse(sort.IntSlice(input)))
	input = append([]int{input[0] + 3}, input...)
	input = append(input, 0)

	// Find the fixed points

	fixedIndices := []int{}
	for i := range input {
		j := math.Min(float64(i+2), float64(len(input)))
		if fixedPoint(input[i:int(j)]) {
			fixedIndices = append(fixedIndices, i)
		}
	}

	miniperms := []int{}
	permap := map[string]int{}
	for i := 0; i < len(fixedIndices)-1; i++ {

		key := fmt.Sprintf("%v-%v", fixedIndices[i], fixedIndices[i+1])
		subtotal := perms(input[fixedIndices[i] : fixedIndices[i+1]+1])
		permap[key] = subtotal
		miniperms = append(miniperms, subtotal)

	}

	for key, item := range permap {
		fmt.Printf("%v: %v\n", key, item)
	}
	result := 1
	for _, item := range miniperms {
		result = result * item
	}

	fmt.Printf("  Answer: %v\n", result)

}

func main() {
	input := helpers.ReadInputInt("input.txt")
	testInput := helpers.ReadInputInt("input.txt")

	puzzle1(input)

	puzzle2(testInput)
}
