package day04

import (
	"fmt"
	"strings"
)

func parseGrid(input string) [][]rune {
	grid := [][]rune{}
	for _, line := range strings.Split(input, "\n") {
		row := []rune{}
		for _, char := range line {
			row = append(row, char)
		}
		grid = append(grid, row)
	}
	return grid
}

func ProcessPart1(input string) int {
	grid := parseGrid(input)
	fmt.Println(grid)
	return 0
}
