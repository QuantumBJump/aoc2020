package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/QuinnStevens/aoc2020/helpers"
)

var (
	north [2]int = [2]int{0, -1}
	east  [2]int = [2]int{1, 0}
	south [2]int = [2]int{0, 1}
	west  [2]int = [2]int{-1, 0}
	debug bool   = false
)

type waypoint struct {
	x int
	y int
}

func (w *waypoint) info() {
	fmt.Printf("Waypoint relative x: %v\n", w.x)
	fmt.Printf("Waypoint relative y: %v\n", w.y)
	fmt.Println()
}

func (w *waypoint) move(facing [2]int, value int) {
	if debug {
		fmt.Printf("Moving waypoint [%v, %v]...\n", facing[0]*value, facing[1]*value)
	}
	w.x += facing[0] * value
	w.y += facing[1] * value
	if debug {
		w.info()
	}
}

func (w *waypoint) rotate(direction bool, value int) {
	if debug {
		dirstr := "left"
		if direction {
			dirstr = "right"
		}
		fmt.Printf("Rotating waypoint %v degrees %v\n", value, dirstr)
	}
	quarterTurns := value / 90
	for i := 0; i < quarterTurns; i++ {
		tmp := w.x
		w.x = w.y
		w.y = tmp
		if direction {
			// Turning right
			w.x = w.x * -1
		} else {
			w.y = w.y * -1
		}
	}
	if debug {
		w.info()
	}
}

type ship struct {
	x        int
	y        int
	facing   [2]int
	waypoint *waypoint
}

func (s *ship) info() {
	fmt.Printf("Ship x: %v\n", s.x)
	fmt.Printf("Ship y: %v\n", s.y)
	fmt.Printf("Ship facing: %v\n", s.facing)
	fmt.Println()
}

func (s *ship) move(facing [2]int, value int) {
	if debug {
		fmt.Printf("Moving [%v, %v]...\n", facing[0]*value, facing[1]*value)
	}
	s.x += facing[0] * value
	s.y += facing[1] * value
	if debug {
		s.info()
	}
}
func (s *ship) movePoint(value int) {
	if debug {
		fmt.Printf("Moving to point ([%v, %v] relative) %v times...\n", s.waypoint.x, s.waypoint.y, value)
	}
	for i := 0; i < value; i++ {
		s.x += s.waypoint.x
		s.y += s.waypoint.y
	}
	if debug {
		s.info()
	}
}

func (s *ship) turn(direction bool, value int) {
	// right = 1, left = 0
	quarterTurns := value / 90
	for i := 0; i < quarterTurns; i++ {
		if direction {
			if debug {
				fmt.Printf("Turning right...\n")
			}
			tmp := s.facing[0]
			s.facing[0] = s.facing[1]
			s.facing[1] = tmp
			s.facing[0] = s.facing[0] * -1
		} else {
			if debug {
				fmt.Printf("Turning left...\n")
			}
			tmp := s.facing[0]
			s.facing[0] = s.facing[1]
			s.facing[1] = tmp
			s.facing[1] = s.facing[1] * -1
		}
		if debug {
			s.info()
		}
	}
}

func puzzle1(input []string) {
	defer helpers.TimeTracker(time.Now(), "Puzzle 1")
	fmt.Println("Puzzle 1:")
	ship := &ship{
		x:      0,
		y:      0,
		facing: east,
	}
	for _, line := range input {
		result := helpers.ParseNamedRegex("^(?P<command>\\w)(?P<value>\\d+)$", line)
		command := result["command"]
		value, err := strconv.Atoi(result["value"])
		if err != nil {
			log.Fatalf("Uh oh: %v", err)
		}
		switch command {
		case "N":
			ship.move(north, value)
		case "E":
			ship.move(east, value)
		case "S":
			ship.move(south, value)
		case "W":
			ship.move(west, value)
		case "F":
			ship.move(ship.facing, value)
		case "L":
			ship.turn(false, value)
		case "R":
			ship.turn(true, value)
		default:
			log.Fatal("Unexpected command")
		}

	}

	fmt.Printf("  Answer: Ship is at [%v, %v]\n\tManhattan distance is %v + %v = %v\n", ship.x, ship.y, math.Abs(float64(ship.x)), math.Abs(float64(ship.y)), math.Abs(float64(ship.x))+math.Abs(float64(ship.y)))
}

func puzzle2(input []string) {
	defer helpers.TimeTracker(time.Now(), "Puzzle 2")
	fmt.Println("Puzzle 2:")
	point := &waypoint{
		x: 10,
		y: -1,
	}
	ship := &ship{
		x:        0,
		y:        0,
		facing:   east,
		waypoint: point,
	}

	for _, line := range input {
		result := helpers.ParseNamedRegex("^(?P<command>\\w)(?P<value>\\d+)$", line)
		command := result["command"]
		value, err := strconv.Atoi(result["value"])
		if err != nil {
			log.Fatalf("Uh oh: %v", err)
		}
		switch command {
		case "N":
			point.move(north, value)
		case "E":
			point.move(east, value)
		case "S":
			point.move(south, value)
		case "W":
			point.move(west, value)
		case "F":
			ship.movePoint(value)
		case "L":
			point.rotate(false, value)
		case "R":
			point.rotate(true, value)
		default:
			log.Fatal("Unexpected command")
		}
	}
	fmt.Printf("  Answer: Ship is at [%v, %v]\n\tManhattan distance is %v + %v = %v\n", ship.x, ship.y, math.Abs(float64(ship.x)), math.Abs(float64(ship.y)), math.Abs(float64(ship.x))+math.Abs(float64(ship.y)))
}

func main() {
	input, err := helpers.ReadInput("input.txt")
	if err != nil {
		log.Fatal("Uh oh")
	}

	puzzle1(input)
	puzzle2(input)
}
