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
