package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/QuinnStevens/aoc2020/helpers"
)

// Instruction represents an assembly language instruction
type Instruction struct {
	command  string
	value    int
	executed int
}

// NewInstruction creates a new Instruction struct
func NewInstruction(cmd string, val int) *Instruction {
	instruction := &Instruction{}
	instruction.command = cmd
	instruction.value = val
	instruction.executed = 0
	return instruction
}

func populateCode(input []string) map[int]*Instruction {
	code := map[int]*Instruction{}
	for i, item := range input {
		var inst *Instruction
		parsed := helpers.ParseNamedRegex("^(?P<command>\\w{3}) (?P<value>[+-]\\d+)$", item)
		val, err := strconv.Atoi(parsed["value"])
		if err != nil {
			log.Fatalf("Error parsing number %v: %v", parsed["value"], err)
		}
		inst = NewInstruction(parsed["command"], val)
		code[i] = inst
	}
	return code
}
func resetCode(code map[int]*Instruction) {
	for _, item := range code {
		item.executed = 0
	}
}

func checkCode(code map[int]*Instruction) (bool, int, int) {
	acc := 0
	for i := 0; i < len(code); i++ {
		inst := code[i]
		// First, check if the instruction has been run before
		if inst.executed > 0 {
			resetCode(code)
			return true, i, acc
		}

		if inst.command == "acc" {
			acc += inst.value
		} else if inst.command == "jmp" {
			i += inst.value
			i--
		}

		inst.executed++
	}
	return false, len(code) - 1, acc
}

func printFirstNCommands(code map[int]*Instruction, lines int) {
	for i := 0; i < lines; i++ {
		fmt.Printf("\t%v", code[i])
	}
	fmt.Printf("\n")
}

func (n *Instruction) toggle() {
	if n.command == "nop" {
		n.command = "jmp"
	} else if n.command == "jmp" {
		n.command = "nop"
	}
	return
}

func puzzle1(input map[int]*Instruction) {
	defer helpers.TimeTracker(time.Now(), "puzzle 1")
	_, line, acc := checkCode(input)
	fmt.Printf("\tCode loops at line %v with accumulator %v\n", line, acc)
}

func puzzle2(input map[int]*Instruction) {
	defer helpers.TimeTracker(time.Now(), "puzzle 2")
	for i := 0; i < len(input); i++ {
		// Check values one at a time
		if input[i].command == "acc" {
			continue // Ignore acc commands
		}
		fmt.Printf("  Toggling line %v...", i)
		input[i].toggle()
		loops, line, acc := checkCode(input)
		if !loops {
			fmt.Printf("Succeeded! Acc: %v\n", acc)
			return
		}
		fmt.Printf("loops at line %v\n", line) // report when it loops
		input[i].toggle()                      // change the instruction back
	}
}

func main() {
	input, err := helpers.ReadInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Puzzle 1
	fmt.Println("Puzzle 1:")
	puzzle1code := populateCode(input)
	puzzle1(puzzle1code)

	// Puzzle 2
	fmt.Println()
	fmt.Println("Puzzle 2:")
	puzzle2code := populateCode(input)
	puzzle2(puzzle2code)
}
