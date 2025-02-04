package day05

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) ([][]int, [][]int) {
	part1 := strings.Split(input, "\n\n")[0]
	part2 := strings.Split(input, "\n\n")[1]
	rules := [][]int{}
	pages := [][]int{}
	for _, line := range strings.Split(part1, "\n") {
		fragments := strings.Split(line, "|")
		page1, _ := strconv.Atoi(fragments[0])
		page2, _ := strconv.Atoi(fragments[1])
		rule := []int{page1, page2}
		rules = append(rules, rule)
	}
	for _, line := range strings.Split(part2, "\n") {
		fragments := strings.Split(line, ",")
		page := []int{}
		for _, fragment := range fragments {
			pageNum, _ := strconv.Atoi(fragment)
			page = append(page, pageNum)
		}
		pages = append(pages, page)
	}
	return rules, pages
}

func ProcessPart1(input string) int {
	rules, pages := parseInput(input)
	orderingRules := map[int][]int{}
	for _, rule := range rules {
		if page, ok := orderingRules[rule[0]]; ok {
			fmt.Println(page)
			page = append(page, rule[1])
			orderingRules[rule[0]] = page
		} else {
			orderingRules[rule[0]] = []int{rule[1]}
		}
	}
	fmt.Println(rules, pages)
	fmt.Println(orderingRules)
	return 0
}
