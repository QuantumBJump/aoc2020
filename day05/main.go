package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/QuinnStevens/aoc2020/helpers"
)

func puzzle1(input []string) {
	var highestID int64 = 0
	highestrow := ""
	highestcol := ""
	var highestrowint int64 = 0
	var highestcolint int64 = 0
	highestlineno := 0

	for i, line := range input {
		// rowstr := ""
		// colstr := ""
		chars := []string{}
		for _, char := range line {
			if char == 'F' || char == 'L' {
				chars = append(chars, "0")
			} else {
				chars = append(chars, "1")
			}
		}
		rowstr := strings.Join(chars[:7], "")
		colstr := strings.Join(chars[7:], "")
		rowint, err := strconv.ParseInt(rowstr, 2, 8)
		if err != nil {
			log.Fatalf("Error converting %v to int: %v", rowstr, err)
		}
		colint, err := strconv.ParseInt(colstr, 2, 8)
		if err != nil {
			log.Fatalf("Error converting %v to int: %v", rowstr, err)
		}

		id := (rowint * 8) + colint
		if id > int64(highestID) {
			highestID = id
			highestrow = rowstr
			highestcol = colstr
			highestrowint = rowint
			highestcolint = colint
			highestlineno = i
		}
		chars = []string{}
	}
	fmt.Printf("\tHighest ID: %d\n", highestID)
	fmt.Printf("\tHighest Row: %s (%d)\n", highestrow, highestrowint)
	fmt.Printf("\tHighest Column: %s (%d)\n", highestcol, highestcolint)
	fmt.Printf("\tHighest Line No: %d\n", highestlineno)
}

func puzzle2(input []string) {
	ids := []int{}

	for _, line := range input {
		// rowstr := ""
		// colstr := ""
		chars := []string{}
		for _, char := range line {
			if char == 'F' || char == 'L' {
				chars = append(chars, "0")
			} else {
				chars = append(chars, "1")
			}
		}
		rowstr := strings.Join(chars[:7], "")
		colstr := strings.Join(chars[7:], "")
		rowint, err := strconv.ParseInt(rowstr, 2, 8)
		if err != nil {
			log.Fatalf("Error converting %v to int: %v", rowstr, err)
		}
		colint, err := strconv.ParseInt(colstr, 2, 8)
		if err != nil {
			log.Fatalf("Error converting %v to int: %v", rowstr, err)
		}

		id := (rowint * 8) + colint
		ids = append(ids, int(id))
		chars = []string{}
	}
	sort.Ints(ids)
	for i, id := range ids {
		if i > 0 {
			// Don't do this for the first one
			if id-ids[i-1] > 1 {
				fmt.Printf("\tYour ID: %d", id-1)
				os.Exit(0)
			}
		}
	}
	fmt.Println("ID not found :(")
}
func main() {
	fmt.Println("Day 5:")
	input, err := helpers.ReadInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	puzzle1(input)
	puzzle2(input)
}
