package main

import (
	"fmt"
	"github.com/patrickjmcd/advent-of-code-2022/helpers"
	"strconv"
	"strings"
)

type pairing struct {
	elf1Start int
	elf2Start int
	elf1End   int
	elf2End   int
}

func (p *pairing) oneElfFullyContainedInOther() bool {
	if p.elf1Start >= p.elf2Start && p.elf1End <= p.elf2End {
		return true
	}
	if p.elf2Start >= p.elf1Start && p.elf2End <= p.elf1End {
		return true
	}
	return false
}

func (p *pairing) hasOverlap() bool {
	// if the start or end is between the other start and end, then there is overlap
	for i := p.elf1Start; i <= p.elf1End; i++ {
		for j := p.elf2Start; j <= p.elf2End; j++ {
			if i == j {
				return true
			}
		}
	}
	return false
}

func newPairing(line string) pairing {
	var p pairing
	individualElves := strings.Split(line, ",")
	elf1 := strings.Split(individualElves[0], "-")
	elf2 := strings.Split(individualElves[1], "-")
	elf1Start, _ := strconv.Atoi(elf1[0])
	elf1End, _ := strconv.Atoi(elf1[1])
	elf2Start, _ := strconv.Atoi(elf2[0])
	elf2End, _ := strconv.Atoi(elf2[1])
	p.elf1Start = elf1Start
	p.elf1End = elf1End
	p.elf2Start = elf2Start
	p.elf2End = elf2End
	return p
}

func part1(filename string) int {
	data, err := helpers.OpenAndReadFile(filename)
	if err != nil {
		panic(err)
	}
	strData := string(data)

	fullyContainedPairings := 0
	//split the file by lines
	lines := strings.Split(strData, "\n")
	// split the lines by commas
	for i := range lines {
		if len(lines[i]) > 0 {
			p := newPairing(lines[i])
			if p.oneElfFullyContainedInOther() {
				fullyContainedPairings += 1
			}
		}
	}
	fmt.Printf("Part 1: %d\n", fullyContainedPairings)
	return fullyContainedPairings
}

func part2(filename string) int {
	data, err := helpers.OpenAndReadFile(filename)
	if err != nil {
		panic(err)
	}
	strData := string(data)

	partiallyContainedPairings := 0
	//split the file by lines
	lines := strings.Split(strData, "\n")
	// split the lines by commas
	for i := range lines {
		if len(lines[i]) > 0 {
			p := newPairing(lines[i])
			if p.hasOverlap() {
				partiallyContainedPairings += 1
			}
		}
	}
	fmt.Printf("Part 2: %d\n", partiallyContainedPairings)
	return partiallyContainedPairings
}

func main() {
	part1("input.txt")
	part2("input.txt")
}
