package day04

import (
	"strings"
)

func getXMASMasks(ints []int) [][][]int {
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

func getMASMasks(ints []int) [][][]int {
	i := ints[0]
	j := ints[1]
	return [][][]int{
		{{i - 1, j - 1}, {i, j}, {i + 1, j + 1}},
		{{i + 1, j + 1}, {i, j}, {i - 1, j - 1}},
		{{i + 1, j - 1}, {i, j}, {i - 1, j + 1}},
		{{i - 1, j + 1}, {i, j}, {i + 1, j - 1}},
	}
}

func isXmas(mask [][]int, grid [][]rune, target string) bool {
	var acc string
	for _, coord := range mask {
		i := coord[0]
		j := coord[1]
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
			return false
		}
		acc += string(grid[i][j])
	}
	return acc == target
}

func calculateXmas(ints []int, grid [][]rune) int {
	masks := getXMASMasks(ints)
	var acc int
	for _, mask := range masks {
		if isXmas(mask, grid, "XMAS") {
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

func ProcessPart2(input string) int {
	var acc int
	grid := parseGrid(input)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			acc += calculateXmas([]int{i, j}, grid)
		}
	}
	return acc
}
