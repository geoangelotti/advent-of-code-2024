package day02

import (
	"strconv"
	"strings"
)

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -1 * x
}

func safeDifferences(levelsDifferences []int) bool {
	for _, levelsDifference := range levelsDifferences {
		if !(abs(levelsDifference) > 0 && abs(levelsDifference) < 4) {
			return false
		}
	}
	return true
}

func safeDirections(levelsDifference []int) bool {
	isPositiveComparison := func(x int) bool {
		return x <= 0
	}
	isNegativeComparison := func(x int) bool {
		return x >= 0
	}
	return allSlice(levelsDifference, isPositiveComparison) || allSlice(levelsDifference, isNegativeComparison)
}

func allSlice(ints []int, comparison func(int) bool) bool {
	for _, i := range ints {
		if comparison(i) {
			return false
		}
	}
	return true
}

func getReports(input string) [][]int {
	reports := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		report := []int{}
		for _, levels := range strings.Split(line, " ") {
			levels, _ := strconv.Atoi(levels)
			report = append(report, levels)
		}
		reports = append(reports, report)
	}
	return reports
}

func ProcessPart1(input string) int {
	acc := 0
	reports := getReports(input)
	for _, report := range reports {
		windowed := []int{}
		for i := 0; i < len(report)-1; i++ {
			windowed = append(windowed, report[i]-report[i+1])
		}
		if safeDifferences(windowed) && safeDirections(windowed) {
			acc += 1
		}
	}
	return acc
}

func ProcessPart2(input string) int {
	return 0
}
