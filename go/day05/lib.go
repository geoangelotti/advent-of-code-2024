package day05

import (
	"fmt"
	"strings"
)

func parseInput(input string) ([]string, [][]string) {
	part1 := strings.Split(input, "\n\n")[0]
	part2 := strings.Split(input, "\n\n")[1]
	rules := strings.Split(part1, "\n")
	pages := [][]string{}
	for _, line := range strings.Split(part2, "\n") {
		fragments := strings.Split(line, ",")
		pages = append(pages, fragments)
	}
	return rules, pages
}

func ProcessPart1(input string) int {
	rules, pages := parseInput(input)
	orderingRules := map[string]bool{}
	for _, rule := range rules {
		orderingRules[rule] = true
	}
	fmt.Println(rules, pages)
	for k, v := range orderingRules {
		fmt.Println(k, v)
	}
	return 0
}
