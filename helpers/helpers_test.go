package helpers

import "testing"

func TestStringSliceAllSameChar(t *testing.T) {
	type args struct {
		a        []string
		expected bool
	}

	var tests = []args{
		{
			a:        []string{"a", "a", "a", "a", "a"},
			expected: true,
		},
		{
			a:        []string{"a", "b", "c", "d", "e"},
			expected: false,
		},
		{
			a:        []string{"a", "a", "a", "a", "b"},
			expected: false,
		},
		{
			a:        []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a"},
			expected: true,
		},
	}

	for _, tt := range tests {
		actual := StringSliceAllSameChar(tt.a)
		if actual != tt.expected {
			t.Errorf("StringSliceAllSameChar(%v) = %v, want %v", tt.a, actual, tt.expected)
		}
	}
}

func TestStringSliceAllUnique(t *testing.T) {

	type args struct {
		a        []string
		expected bool
	}

	var tests = []args{
		{
			a:        []string{"a", "b", "c", "d", "e"},
			expected: true,
		},
		{
			a:        []string{"a", "a", "a", "a", "b"},
			expected: false,
		},
		{
			a:        []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a"},
			expected: false,
		},
	}

	for _, tt := range tests {
		actual := StringSliceAllUnique(tt.a)
		if actual != tt.expected {
			t.Errorf("StringSliceAllUnique(%v) = %v, want %v", tt.a, actual, tt.expected)
		}
	}

}
