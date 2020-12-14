package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/QuinnStevens/aoc2020/helpers"
)

type bag struct {
	style string
	count int
	sub   []bag
}

func newBag(style string, count int) bag {
	bag := bag{
		style: style,
		count: count,
		sub:   []bag{},
	}
	return bag
}

func (b *bag) CanContain(bags map[string]bag, target string) bool {
	for _, sb := range b.sub {
		if sb.style == target {
			return true
		}
		t := bags[sb.style]
		result := t.CanContain(bags, target)
		if result {
			return true
		}
	}
	return false
}

func (b bag) ContainsTotal(bags map[string]bag) int {
	total := 0
	for _, sb := range b.sub {
		total += sb.count
		for i := 0; i < sb.count; i++ {
			t := bags[sb.style]
			sbTotal := t.ContainsTotal(bags)
			total += sbTotal
		}
	}
	return total
}

func puzzle1(input []string) {
	allowed := map[string]bag{}
	for _, line := range input {
		result := helpers.ParseNamedRegex("^(?P<container>[\\w\\s]*) bags contain (?P<contains>.*)\\.$", line)
		container := result["container"]
		b := bag{}
		b.style = container
		contains := []string{}
		for _, item := range strings.Split(result["contains"], ", ") {
			contains = append(contains, item)
		}
		for _, item := range contains {
			re, err := regexp.Compile("^(\\d+) ([\\w\\s]*) bags?\\.?$")
			if err != nil {
				log.Fatalf("Error parsing re: %v", err)
			}
			match := re.FindStringSubmatch(item)
			if match != nil {
				// Only do this if there are bags allowed
				subBag := bag{}
				subBag.count, err = strconv.Atoi(match[1])
				if err != nil {
					log.Fatalf("Error parsing number %s: %v", match[1], err)
				}
				subBag.style = match[2]
				b.sub = append(b.sub, subBag)
			}
		}
		allowed[container] = b
	}

	total := 0
	for _, value := range allowed {
		// for _, content := range value.sub {
		// 	fmt.Printf("  %s: %d\n", content.style, content.count)
		// }
		contains := value.CanContain(allowed, "shiny gold")
		if contains {
			total++
		}
	}
	fmt.Printf("Total bags which can contain gold: %d\n", total)
	fmt.Printf("Total number of bags in a gold bag: %d\n", allowed["shiny gold"].ContainsTotal(allowed))
}

func main() {
	input, err := helpers.ReadInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	puzzle1(input)
}
