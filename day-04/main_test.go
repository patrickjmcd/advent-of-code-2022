package main

import "testing"

type args struct {
	p           pairing
	contained   bool
	overlapping bool
}

var tests = []args{
	// elf 2 completely contained
	{
		p: pairing{
			elf1Start: 1,
			elf1End:   4,
			elf2Start: 2,
			elf2End:   3,
		},
		contained:   true,
		overlapping: true,
	},
	// elf 1 completely contained
	{
		p: pairing{
			elf1Start: 2,
			elf1End:   3,
			elf2Start: 1,
			elf2End:   4,
		},
		contained:   true,
		overlapping: true,
	},
	// has same end
	{
		p: pairing{
			elf1Start: 1,
			elf1End:   4,
			elf2Start: 2,
			elf2End:   4,
		},
		contained:   true,
		overlapping: true,
	},
	// has same start
	{
		p: pairing{
			elf1Start: 1,
			elf1End:   4,
			elf2Start: 1,
			elf2End:   3,
		},
		contained:   true,
		overlapping: true,
	},
	// not touching
	{
		p: pairing{
			elf1Start: 1,
			elf1End:   2,
			elf2Start: 3,
			elf2End:   4,
		},
		contained:   false,
		overlapping: false,
	},
	// not contained
	{
		p: pairing{
			elf1Start: 1,
			elf1End:   2,
			elf2Start: 2,
			elf2End:   3,
		},
		contained:   false,
		overlapping: true,
	},
	{
		p: pairing{
			elf1Start: 2,
			elf1End:   3,
			elf2Start: 1,
			elf2End:   2,
		},
		contained:   false,
		overlapping: true,
	},
	{
		p: pairing{
			elf1Start: 1,
			elf1End:   2,
			elf2Start: 1,
			elf2End:   2,
		},
		contained:   true,
		overlapping: true,
	},
}

func TestPairingCompletelyContained(t *testing.T) {
	for _, test := range tests {
		c := test.p.oneElfFullyContainedInOther()
		if c != test.contained {
			t.Errorf("Expected %t, got %t", test.contained, c)
		}
	}

}

func TestPairingOverlap(t *testing.T) {
	for _, test := range tests {
		c := test.p.hasOverlap()
		if c != test.overlapping {
			t.Errorf("Expected %t, got %t for %+v", test.overlapping, c, test.p)
		}
	}
}

func TestPart1(t *testing.T) {
	score := part1("sample.txt")
	if score != 2 {
		t.Errorf("Expected 2, got %d", score)
	}
}

func TestPart2(t *testing.T) {
	score := part2("sample.txt")
	if score != 4 {
		t.Errorf("Expected 4, got %d", score)
	}
}
