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
		isIncreasing := report[1] > report[0]
		isSafe := true
		for i := 1; i < len(report); i++ {
			if report[i] == report[i-1] {
				isSafe = false
				break
			}
			if isIncreasing != (report[i] > report[i-1]) {
				isSafe = false
				break
			}
			if abs(report[i]-report[i-1]) > 3 {
				isSafe = false
				break
			}
		}
		if isSafe {
			acc += 1
		}
	}
	return acc
}

func ProcessPart2(input string) int {
	return 0
}
