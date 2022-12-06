package main

import (
	"github.com/patrickjmcd/advent-of-code-2022/helpers"
	"testing"
)

func TestGetCharsForCommsCheck(t *testing.T) {
	type args struct {
		input    []string
		expected []string
	}
	var tests = []args{
		{
			input:    []string{"", "", "", "d", "e"},
			expected: []string{"d", "e", "", ""},
		},
		{
			input:    []string{"", "", "c", "d", "e", "f", "g", "h", "i"},
			expected: []string{"f", "g", "h", "i"},
		},
		{
			input:    []string{"a", "b", "c", "d"},
			expected: []string{"a", "b", "c", "d"},
		},
	}
	for _, tt := range tests {
		actual := getCharsForCommsCheck(tt.input)
		if !helpers.StringSliceEqual(actual, tt.expected) {
			t.Errorf("getCharsForCommsCheck(%s) = %v, want %v", tt.input, actual, tt.expected)
		}
	}
}

func TestFindStartOfCommunication(t *testing.T) {
	type args struct {
		input    string
		expected int
	}

	var tests = []args{
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 5,
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 6,
		},
		{
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 10,
		},
		{
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 11,
		},
	}
	for _, tt := range tests {
		clearProcessed()
		actual := findStartOfCommunication(tt.input)
		if actual != tt.expected {
			t.Errorf("findStartOfCommunication(%s) = %v, want %v", tt.input, actual, tt.expected)
		}
	}
}

func TestFindStartOfMessage(t *testing.T) {
	type args struct {
		input    string
		expected int
	}
	var tests = []args{
		{
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: 19,
		},
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 23,
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 23,
		},
		{
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 29,
		},
		{
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 26,
		},
	}
	for _, tt := range tests {
		clearProcessed()
		actual := findStartOfMessage(tt.input)
		if actual != tt.expected {
			t.Errorf("findStartOfMessage(%s) = %v, want %v", tt.input, actual, tt.expected)
		}
	}
}
