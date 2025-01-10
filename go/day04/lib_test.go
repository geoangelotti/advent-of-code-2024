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

const INPUT2 = `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`

func TestProcessPart1(t *testing.T) {
	expected := 18
	answer := day04.ProcessPart1(INPUT1)
	if answer != expected {
		t.Errorf("Expected %d got %d", expected, answer)
	}
}

func TestProcessPart2(t *testing.T) {
	expected := 9
	answer := day04.ProcessPart2(INPUT1)
	if answer != expected {
		t.Errorf("Expected %d got %d", expected, answer)
	}
}
