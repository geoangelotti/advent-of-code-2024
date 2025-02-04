package day05

import (
	"strconv"
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

func filterPages(pages []string, rules map[string]bool) bool {
	return false
}

func ProcessPart1(input string) int {
	var acc int
	rules, pages := parseInput(input)
	orderingRules := map[string]bool{}
	for _, rule := range rules {
		orderingRules[rule] = true
	}

	for _, line := range pages {
		if filterPages(line, orderingRules) {
			middlePage, _ := strconv.Atoi(line[len(line)/2])
			acc += middlePage
		}
	}
	return acc
}
