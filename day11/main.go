package main

import (
	"fmt"
	"log"
	"time"

	"github.com/QuinnStevens/aoc2020/helpers"
)

type space struct {
	value     string
	lastValue string
	nextValue string
	x         int
	y         int
}

func (s *space) calc(input [][]*space) {
	debug := false
	// if s.y == 9 && s.x == 8 {
	// 	debug = true
	// }
	nearbyOccupied := 0
	for y := s.y - 1; y < s.y+2; y++ {
		if y < 0 || y >= len(input) {
			// invalid, ignore
			continue
		}
		for x := s.x - 1; x < s.x+2; x++ {
			if x == s.x && y == s.y {
				// Don't count yourself
				continue
			}
			if x < 0 || x >= len(input[y]) {
				// invalid, ignore
				continue
			}
			if debug {
				fmt.Printf("Checking space [%v, %v], value %v\n", x, y, input[y][x].value)
			}
			if input[y][x].value == "#" {
				nearbyOccupied++
			}
		}
	}
	if debug {
		fmt.Printf("Nearby occupied seats: %v\n", nearbyOccupied)
	}
	if s.value == "L" && nearbyOccupied == 0 {
		// Seat becomes occupied
		s.nextValue = "#"
		return
	}
	if s.value == "#" && nearbyOccupied >= 4 {
		// Seat becomes empty
		s.nextValue = "L"
		return
	}
	s.nextValue = s.value
}

func (s *space) calc2(input [][]*space) {
	debug := false
	nearbyOccupied := 0
	directions := [][2]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	// diag up-left

	for _, pair := range directions {
		x := s.x
		y := s.y
		for {
			x += pair[0]
			y += pair[1]
			if x < 0 || y < 0 || y >= len(input) || x >= len(input[0]) {
				// out of bounds
				break
			}
			if input[y][x].value == "." {
				// Not a chair, keep looking
				continue
			} else {
				if debug {
					fmt.Printf("Found chair at [%v, %v]: %v\n", x, y, input[y][x].value)
				}
				if input[y][x].value == "#" {
					nearbyOccupied++
				}
				break
			}
		}

	}
	if s.value == "L" && nearbyOccupied == 0 {
		// Seat becomes occupied
		s.nextValue = "#"
		return
	}
	if s.value == "#" && nearbyOccupied >= 5 {
		// Seat becomes empty
		s.nextValue = "L"
		return
	}
	s.nextValue = s.value
}

func (s *space) update() {
	s.lastValue = s.value
	s.value = s.nextValue
	s.nextValue = ""
}

func printgrid(input [][]*space) {
	fmt.Println()
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			fmt.Printf("%v", input[y][x].value)
		}
		fmt.Printf("\n")
	}
}

func updategrid(input [][]*space) bool {
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			input[y][x].calc(input)
		}
	}

	if stagnant(input) {
		return true
	}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			input[y][x].update()
		}
	}
	return false
}

func stagnant(input [][]*space) bool {
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x].value != input[y][x].nextValue {
				return false
			}
		}
	}
	return true
}

func loop(input [][]*space) {
	finished := false
	debug := false
	for !finished {
		if debug {
			printgrid(input)
		}
		finished = updategrid(input)
	}
}
func puzzle1(input []string) {
	defer helpers.TimeTracker(time.Now(), "Puzzle 1")
	fmt.Println("Puzzle 1:")
	floor := [][]*space{}
	for y, line := range input {
		row := []*space{}
		for x, char := range line {
			sp := &space{
				value:     string(char),
				lastValue: "",
				nextValue: "",
				x:         x,
				y:         y,
			}
			row = append(row, sp)
		}
		floor = append(floor, row)
	}
	loop(floor)

	occupied := 0
	for y := 0; y < len(floor); y++ {
		for x := 0; x < len(floor[y]); x++ {
			if floor[y][x].value == "#" {
				occupied++
			}
		}
	}

	fmt.Printf("  Answer: %v\n", occupied)
}

func updategrid2(input [][]*space) bool {
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			input[y][x].calc2(input)
		}
	}

	if stagnant(input) {
		return true
	}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			input[y][x].update()
		}
	}
	return false
}

func loop2(input [][]*space) {
	finished := false
	debug := false
	for !finished {
		if debug {
			printgrid(input)
		}
		finished = updategrid2(input)
	}
}

func puzzle2(input []string) {
	defer helpers.TimeTracker(time.Now(), "Puzzle 2")
	fmt.Println("Puzzle 2:")
	floor := [][]*space{}
	for y, line := range input {
		row := []*space{}
		for x, char := range line {
			sp := &space{
				value:     string(char),
				lastValue: "",
				nextValue: "",
				x:         x,
				y:         y,
			}
			row = append(row, sp)
		}
		floor = append(floor, row)
	}
	loop2(floor)

	occupied := 0
	for y := 0; y < len(floor); y++ {
		for x := 0; x < len(floor[y]); x++ {
			if floor[y][x].value == "#" {
				occupied++
			}
		}
	}

	fmt.Printf("  Answer: %v\n", occupied)
}

func main() {
	input, err := helpers.ReadInput("input.txt")
	if err != nil {
		log.Fatal("Oh no")
	}

	puzzle1(input)
	puzzle2(input)
}
