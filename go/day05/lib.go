package day05

import (
	"strconv"
	"strings"
)

func parseInput(input string) (map[string]bool, [][]string) {
	part1 := strings.Split(input, "\n\n")[0]
	part2 := strings.Split(input, "\n\n")[1]
	pages := [][]string{}
	rules := map[string]bool{}
	for _, rule := range strings.Split(part1, "\n") {
		rules[rule] = true
	}
	for _, line := range strings.Split(part2, "\n") {
		fragments := strings.Split(line, ",")
		pages = append(pages, fragments)
	}
	return rules, pages
}

func filterIncorrectPages(pages []string, rules map[string]bool) []int {
	incorrectPages := []int{}
	for i, page := range pages {
		for j := 0; j < i; j++ {
			composite := pages[j] + "|" + page
			if !rules[composite] {
				incorrectPages = append(incorrectPages, i)
			}
		}
	}
	return incorrectPages
}

func filterPages(pages []string, rules map[string]bool) bool {
	return len(filterIncorrectPages(pages, rules)) <= 0
}

func rearrangePages(pages []string, indexes []int, rules map[string]bool) []string {
	return pages
}

func ProcessPart1(input string) int {
	var acc int
	rules, pages := parseInput(input)
	for _, line := range pages {
		if filterPages(line, rules) {
			middlePage, _ := strconv.Atoi(line[len(line)/2])
			acc += middlePage
		}
	}
	return acc
}

func ProcessPart2(input string) int {
	var acc int
	rules, pages := parseInput(input)
	for _, line := range pages {
		if incorrectPages := filterIncorrectPages(line, rules); len(incorrectPages) > 0 {
			rearranged := rearrangePages(line, incorrectPages, rules)
			middlePage, _ := strconv.Atoi(rearranged[len(rearranged)/2])
			acc += middlePage
		}
	}
	return acc
}
