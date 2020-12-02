package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readLines(path string) ([]int, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return lines, err
		}
		lines = append(lines, num)
	}
	return lines, scanner.Err()
}
func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("Error reading file!: %v", err)
	}

	for i, line := range lines {
		compare := lines[i:]
		for j, item := range compare {
			sum := line + item
			if sum > 2020 {
				continue
			} else {
				compare2 := compare[j:]
				for _, thing := range compare2 {
					sum := line + item + thing
					if sum == 2020 {
						println("product: %d", line*item*thing)
					}
				}
			}
		}
	}
}
