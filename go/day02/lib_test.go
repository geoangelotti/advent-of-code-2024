package day02_test

import (
	"day02"
	"testing"
)

const INPUT = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestProcessPart1(t *testing.T) {
	expected := 2
	answer := day02.ProcessPart1(INPUT)
	if answer != expected {
		t.Errorf("Expected %d got %d", expected, answer)
	}
}

func TestProcessPart2(t *testing.T) {
	expected := 4
	answer := day02.ProcessPart2(INPUT)
	if answer != expected {
		t.Errorf("Expected %d got %d", expected, answer)
	}
}
