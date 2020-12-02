package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("Error reading file!: %v", err)
	}

	var item1 int
	var item2 int
	for i, line := range lines {
		item1, err = strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Unable to read line %v", line)
		}
		compare := lines[i:]
		for _, item := range compare {
			item2, err = strconv.Atoi(item)
			if err != nil {
				log.Fatalf("Unable to read line %v", line)
			}
			sum := item1 + item2
			if sum == 2020 {
				println(item1 * item2)
				os.Exit(0)
			}
		}
	}
}
