package main

import "testing"

func TestGetLetterPriority(t *testing.T) {
	tests := []struct {
		name     string
		letter   string
		expected int
	}{
		{
			name:     "Lowercase",
			letter:   "a",
			expected: 1,
		},
		{
			name:     "Uppercase",
			letter:   "A",
			expected: 27,
		},
		{
			name:     "Space",
			letter:   " ",
			expected: -1,
		},
		{
			name:     "Non-ASCII",
			letter:   "😀",
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := getLetterPriority(test.letter)
			if actual != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, actual)
			}
		})
	}
}

func TestFindCommonLetters(t *testing.T) {
	word1 := "vJrwpWtwJgWr"
	word2 := "hcsFMMfFFhFp"
	expected := []string{"p"}
	commonLetters := findCommonLetters(word1, word2)
	for i := range commonLetters {
		if commonLetters[i] != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], commonLetters[i])
		}
	}
}

func TestPart1(t *testing.T) {
	score := part1("sample.txt")
	if score != 157 {
		t.Errorf("Expected 157, got %d", score)
	}
}

func TestPart2(t *testing.T) {
	score := part2("sample.txt")
	if score != 70 {
		t.Errorf("Expected 70, got %d", score)
	}
}
