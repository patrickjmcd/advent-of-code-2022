package main

import (
	"fmt"
	"github.com/patrickjmcd/advent-of-code-2022/helpers"
	"regexp"
	"strings"
	"testing"
)

func TestParseCommand(t *testing.T) {
	type args struct {
		line     string
		expected command
	}

	var tests = []args{
		{
			line: "move 1 from 2 to 1",
			expected: command{
				quantity: 1,
				from:     2,
				to:       1,
			},
		},
		{
			line: "move 3 from 1 to 3",
			expected: command{
				quantity: 3,
				from:     1,
				to:       3,
			},
		},
		{
			line: "move 2 from 2 to 1",
			expected: command{
				quantity: 2,
				from:     2,
				to:       1,
			},
		},
		{
			line: "move 1 from 1 to 2",
			expected: command{
				quantity: 1,
				from:     1,
				to:       2,
			},
		},
	}

	for _, tt := range tests {
		c := parseCommand(tt.line)
		if c != tt.expected {
			t.Errorf("parseCommand(%s) = %v, want %v", tt.line, c, tt.expected)
		}

	}
}

func TestSplitting(t *testing.T) {
	inp := " 1   2   3   4   5   6   7   8   9 "
	r := regexp.MustCompile(`\d+`)
	matches := r.FindAllString(inp, -1)

	if len(matches) != 9 {
		t.Errorf("Expected 9, got %d (%v)", len(matches), matches)
	}
}

func TestParseDrawing(t *testing.T) {
	data, _ := helpers.OpenAndReadFile("sample.txt")
	strData := string(data)
	split := strings.Split(strData, "\n\n")

	containers, _ := parseDrawing(split[0])
	fmt.Printf("%+v\n", containers)
	if len(containers) != 6 {
		t.Errorf("Expected 6 containers, got %d (%+v)", len(containers), containers)
	}

}

func TestMakeStacksFromContainer(t *testing.T) {
	data, _ := helpers.OpenAndReadFile("sample.txt")
	strData := string(data)
	split := strings.Split(strData, "\n\n")

	containers, _ := parseDrawing(split[0])
	stacks := makeStacksFromContainers(containers)
	fmt.Printf("%+v\n", stacks)
	if len(stacks) != 3 {
		t.Errorf("Expected 3 stacks, got %d (%+v)", len(stacks), stacks)
	}
}

func TestMoveFromTo(t *testing.T) {
	data := map[int][]string{
		1: {"N", "Z"},
		2: {"D", "C", "M"},
		3: {"P"},
	}
	expected1 := map[int][]string{
		1: {"D", "N", "Z"},
		2: {"C", "M"},
		3: {"P"},
	}

	moveFromTo(2, 1, 1, &data, true)
	for k, v := range data {
		if !helpers.StringSliceEqual(v, expected1[k]) {
			t.Errorf("Expected %v, got %v", expected1[k], v)
		}
	}

	expected2 := map[int][]string{
		1: {},
		2: {"C", "M"},
		3: {"Z", "N", "D", "P"},
	}
	moveFromTo(1, 3, 3, &data, true)
	for k, v := range data {
		if !helpers.StringSliceEqual(v, expected2[k]) {
			t.Errorf("Expected %v, got %v", expected2[k], v)
		}
	}

	expected3 := map[int][]string{
		1: {"M", "C"},
		2: {},
		3: {"Z", "N", "D", "P"},
	}
	moveFromTo(2, 1, 2, &data, true)
	for k, v := range data {
		if !helpers.StringSliceEqual(v, expected3[k]) {
			t.Errorf("Expected %v, got %v", expected3[k], v)
		}
	}

}

func TestPart1(t *testing.T) {
	result, st := part1("sample.txt", true)
	if result != "CMZ" {
		t.Errorf("Expected CMZ, got %s (%+v)", result, st)
	}

	result2, st2 := part1("sample.txt", false)
	if result2 != "MCD" {
		t.Errorf("Expected MCD, got %s (%+v)", result2, st2)
	}
}
