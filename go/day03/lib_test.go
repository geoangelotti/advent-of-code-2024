package day03_test

import (
	"day03"
	"testing"
)

const INPUT = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

func TestProcessPart1(t *testing.T) {
	expected := 161
	answer := day03.ProcessPart1(INPUT)
	if answer != expected {
		t.Errorf("Expected %d got %d", expected, answer)
	}
}
