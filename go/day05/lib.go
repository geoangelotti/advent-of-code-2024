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
	fmt.Println(rules, pages)
	return 0
}
