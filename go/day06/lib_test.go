package day06_test

import (
	"day06"
	"testing"
)

const INPUT = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestProcessPart1(t *testing.T) {
	expected := 41
	answer := day06.ProcessPart1(INPUT)
	if answer != expected {
		t.Errorf("Expected %d got %d", expected, answer)
	}
}
