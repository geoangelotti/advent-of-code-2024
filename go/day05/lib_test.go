package day05_test

import (
	"day05"
	"testing"
)

const INPUT = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestProcessPart1(t *testing.T) {
	expected := 143
	answer := day05.ProcessPart1(INPUT)
	if answer != expected {
		t.Errorf("Expected %d got %d", expected, answer)
	}
}
