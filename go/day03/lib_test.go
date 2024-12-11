package day03_test

import (
	"day03"
	"testing"
)

const INPUT1 = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
const INPUT2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestProcessPart1(t *testing.T) {
	expected := 161
	answer := day03.ProcessPart1(INPUT1)
	if answer != expected {
		t.Errorf("Expected %d got %d", expected, answer)
	}
}

func TestProcessPart2(t *testing.T) {
	expected := 48
	answer := day03.ProcessPart2(INPUT2)
	if answer != expected {
		t.Errorf("Expected %d got %d", expected, answer)
	}
}
