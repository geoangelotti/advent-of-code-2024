package day07

import (
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

func generateCombinations(n int, current string, result *[]string) {
	if n == 0 {
		*result = append(*result, current)
		return
	}

	generateCombinations(n-1, current+"+", result)
	generateCombinations(n-1, current+"*", result)
}

func getCombinations(n int) []string {
	result := []string{}
	generateCombinations(n, "", &result)
	return result
}

func canBeCalibrated(target int, equations []int) bool {
	operations := getCombinations(len(equations) - 1)
	for _, operation := range operations {
		acc := equations[0]
		for i, equation := range equations[1:] {
			if operation[i] == '+' {
				acc += equation
			} else {
				acc *= equation
			}
		}
		if acc == target {
			return true
		}
	}
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
	return acc
}

func ProcessPart2(input string) int {
	return 0
}
