package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/QuinnStevens/aoc2020/helpers"
)

func puzzle1(input []string) {
	defer helpers.TimeTracker(time.Now(), "Puzzle 1")
	fmt.Println("Puzzle 1:")

	earliest, err := strconv.Atoi(input[0])
	if err != nil {
		log.Fatalf("Failed to parse int %v: %v", input[0], err)
	}
	timetable := strings.Split(input[1], ",")
	running := []int{}
	for _, bus := range timetable {
		if bus != "x" {
			num, err := strconv.Atoi(bus)
			if err != nil {
				log.Fatalf("Failed to parse int %v: %v", input[0], err)
			}
			running = append(running, num)
		}
	}
	waitTimes := []int{}
	for _, bus := range running {
		waitTimes = append(waitTimes, (earliest/bus+1)*bus-earliest)
	}

	lowest := waitTimes[0]
	lowestIndex := 0

	for i, time := range waitTimes {
		if time < lowest {
			lowest = time
			lowestIndex = i
		}
	}
	answer := lowest * running[lowestIndex]
	fmt.Printf("  Answer: %v\n", answer)
}

func lcm(a, b int) int {
	gcd := func(a, b int) int {
		for b != 0 {
			t := b
			b = a % b
			a = t
		}
		return a
	}
	return a * b / gcd(a, b)
}

func puzzle2(input []string) {
	defer helpers.TimeTracker(time.Now(), "Puzzle 2")
	fmt.Println("Puzzle 2:")
	timetable := strings.Split(input[1], ",")
	firstBus, err := strconv.Atoi(timetable[0])
	if err != nil {
		log.Fatalf("Error parsing number %v: %v", timetable[0], err)
	}
	t := firstBus     // equal to, e.g., 7
	delta := firstBus // we start off by jumping 7 at a time
	for i, id := range timetable {
		// look at the next number along, e.g. 13
		if id == "x" {
			continue
		}
		num, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal("Oh dear")
		}
		for {
			// Check if t+1 is divisible by 13
			if (i+t)%num == 0 {
				// if it is, we have some t where t%7, t+1%13 == 0
				break
			}
			t += delta
		}
		// now we've found t%7, t+1%13 == 0, we know that the next number to be divisible by both will be the next
		// lowest common multiple of 77 & 7
		delta = lcm(delta, num)
	}

	fmt.Printf("  Answer: t=%v\n", t)
}

func main() {
	input, err := helpers.ReadInput("input.txt")
	if err != nil {
		log.Fatal("Oh dear")
	}

	puzzle1(input)
	puzzle2(input)
}
