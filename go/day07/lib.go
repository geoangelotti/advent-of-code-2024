package day07

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) map[int][]int {
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

func canBeCalibrated(target int, equations []int) bool {
	return false
}

func ProcessPart1(input string) int {
	var acc int
	equations := parseInput(input)
	for key, value := range equations {
		if canBeCalibrated(key, value) {
			acc += key
		}
	}
	fmt.Println(equations)
	return acc
}
