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

func generateCombinations(n int, current string, result *[]string, operations []string) {
	if n == 0 {
		*result = append(*result, current)
		return
	}
	for _, operation := range operations {
		generateCombinations(n-1, current+operation, result, operations)
	}

}

func getCombinations(n int, operations []string) []string {
	result := []string{}
	generateCombinations(n, "", &result, operations)
	return result
}

func canBeCalibrated(target int, numbers []int, allowedOperations []string) bool {
	operations := getCombinations(len(numbers)-1, allowedOperations)
	for _, operation := range operations {
		acc := numbers[0]
		for i, number := range numbers[1:] {
			if operation[i] == '+' {
				acc += number
			} else if operation[i] == '*' {
				acc *= number
			} else {
				concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", acc, number))
				acc = concatenated
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
		if canBeCalibrated(key, value, []string{"+", "*"}) {
			acc += key
		}
	}
	return acc
}

func ProcessPart2(input string) int {
	var acc int
	equations := parseInput(input)
	for key, value := range equations {
		if canBeCalibrated(key, value, []string{"+", "*", " "}) {
			acc += key
		}
	}
	return acc
}
