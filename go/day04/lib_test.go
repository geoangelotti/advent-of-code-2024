package day04_test

import (
	"day04"
	"testing"
)

const INPUT1 = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestProcessPart1(t *testing.T) {
	expected := 18
	answer := day04.ProcessPart1(INPUT1)
	if answer != expected {
		t.Errorf("Expected %d got %d", expected, answer)
	}
}
