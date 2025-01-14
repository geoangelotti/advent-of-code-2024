package day06

import (
	"fmt"
	"strings"
)

type Point []int

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

func parseGrid(input string) [][]rune {
	var grid [][]rune
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}
	return grid
}

func findStart(grid [][]rune) Point {
	for y, row := range grid {
		for x, cell := range row {
			if cell == '^' {
				return Point{x, y}
			}
		}
	}
	return Point{-1, -1}
}

func (s Point) String() string {
	return fmt.Sprintf("%d,%d", s[0], s[1])
}

func nextPoint(point Point, direction int) Point {
	i := point[0]
	j := point[1]
	switch direction {
	case UP:
		return Point{i, j - 1}
	case RIGHT:
		return Point{i + 1, j}
	case DOWN:
		return Point{i, j + 1}
	case LEFT:
		return Point{i - 1, j}
	default:
		return Point{-1, -1}
	}
}

func turnRight(direction int) int {
	return (direction + 1) % 4
}

func ProcessPart1(input string) int {
	var grid = parseGrid(input)
	var start = findStart(grid)
	current := start
	var seen = make(map[string]bool)
	seen[current.String()] = true
	direction := UP
	next := nextPoint(current, direction)
	for next[0] >= 0 && next[1] >= 0 && next[0] < len(grid[0]) && next[1] < len(grid) {
		if grid[next[1]][next[0]] == '#' {
			direction = turnRight(direction)
		} else {
			current = next
			seen[current.String()] = true
		}
		next = nextPoint(current, direction)

	}
	return len(seen)
}

func ProcessPart2(input string) int {
	var acc int
	return acc
}
