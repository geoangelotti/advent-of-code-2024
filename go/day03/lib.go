package day03

import (
	"regexp"
	"strconv"
)

func ProcessPart1(input string) int {
	var acc int
	pattern := regexp.MustCompile(`mul\((?P<a>[0-9]{1,3}),(?P<b>[0-9]{1,3})\)`)
	results := pattern.FindAllStringSubmatch(input, -1)
	for _, r := range results {
		a, _ := strconv.Atoi(r[1])
		b, _ := strconv.Atoi(r[2])
		acc += a * b
	}
	return acc
}

func ProcessPart2(input string) int {
	var acc int
	pattern := regexp.MustCompile(`(mul\((?P<a>[0-9]{1,3}),(?P<b>[0-9]{1,3})\))|do\(\)|don't\(\)`)
	matches := pattern.FindAllStringSubmatch(input, -1)
	do := true
	for _, match := range matches {
		switch match[0] {
		case "do()":
			do = true
		case "don't()":
			do = false
		default:
			if do {
				a, _ := strconv.Atoi(match[2])
				b, _ := strconv.Atoi(match[3])
				acc += a * b
			}
		}
	}
	return acc
}
