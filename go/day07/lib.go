package day07

import (
	"fmt"
	"strconv"
	"strings"
)

func parse_input(input string) map[int][]int {
	equations := make(map[int][]int)
	for _, line := range strings.Split(input, "\n") {
		fragment := strings.Split(line, ":")
		numbers := strings.Split(fragment[1], " ")
		nums := []int{}
		for _, number := range numbers {
			if number == "" {
				continue
			}
			num, _ := strconv.Atoi(number)
			nums = append(nums, num)
		}
		target, _ := strconv.Atoi(fragment[0])
		equations[target] = nums
	}
	return equations
}

func ProcessPart1(input string) int {
	equations := parse_input(input)
	fmt.Println(equations)
	return 0
}
