package day04

import (
	"strings"
)

func getMasks(ints []int) [][][]int {
	i := ints[0]
	j := ints[1]
	return [][][]int{
		{{i, j}, {i + 1, j}, {i + 2, j}, {i + 3, j}},
		{{i, j}, {i - 1, j}, {i - 2, j}, {i - 3, j}},
		{{i, j}, {i, j + 1}, {i, j + 2}, {i, j + 3}},
		{{i, j}, {i, j - 1}, {i, j - 2}, {i, j - 3}},
		{{i, j}, {i + 1, j + 1}, {i + 2, j + 2}, {i + 3, j + 3}},
		{{i, j}, {i + 1, j - 1}, {i + 2, j - 2}, {i + 3, j - 3}},
		{{i, j}, {i - 1, j + 1}, {i - 2, j + 2}, {i - 3, j + 3}},
		{{i, j}, {i - 1, j - 1}, {i - 2, j - 2}, {i - 3, j - 3}},
	}
}

func isXmas(mask [][]int, grid [][]rune) (is bool) {
	defer func() {
		if r := recover(); r != nil {
			is = false
		}
	}()
	var acc string
	for _, coord := range mask {
		i := coord[0]
		j := coord[1]
		acc += string(grid[i][j])
	}
	if acc == "XMAS" {
		return true
	}
	return
}

func calculateXmas(ints []int, grid [][]rune) int {
	masks := getMasks(ints)
	var acc int
	for _, mask := range masks {
		if isXmas(mask, grid) {
			acc += 1
		}
	}
	return acc
}

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
	var acc int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			acc += calculateXmas([]int{i, j}, grid)
		}
	}
	return acc
}
