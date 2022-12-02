package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	score := part1("sample.txt")
	if score != 15 {
		t.Errorf("Expected 15, got %d", score)
	}
}

func TestPart2(t *testing.T) {
	score := part2("sample.txt")
	if score != 12 {
		t.Errorf("Expected 12, got %d", score)
	}
}
