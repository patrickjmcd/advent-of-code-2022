package main

import (
	"fmt"
	"github.com/patrickjmcd/advent-of-code-2022/helpers"
)

var consecutiveCharsComms = 4
var consecutiveCharsMessage = 14
var processed = make([]string, consecutiveCharsMessage)

func clearProcessed() {
	for i := range processed {
		processed[i] = ""
	}
}

func getCharsForCommsCheck(strArr []string) []string {
	var chars = make([]string, consecutiveCharsComms)
	if getNonEmptyLength(strArr) >= consecutiveCharsComms {
		return strArr[len(strArr)-consecutiveCharsComms:]
	}

	filled := 0
	for i := 0; i < len(strArr); i++ {
		if strArr[i] != "" {
			chars[filled] = strArr[i]
			filled++
			if filled == consecutiveCharsComms {
				return chars
			}
		}
	}
	return chars
}

func getNonEmptyLength(arr []string) int {
	var count int
	for _, s := range arr {
		if s != "" {
			count++
		}
	}
	return count
}

func processNewChar(ch uint8) {
	for i := 0; i < len(processed)-1; i++ {
		processed[i] = processed[i+1]
	}
	processed[len(processed)-1] = string(ch)

}

func findStartOfCommunication(input string) int {
	for idx := range input {
		char := input[idx]
		processNewChar(char)
		checkChars := getCharsForCommsCheck(processed)
		if getNonEmptyLength(checkChars) >= consecutiveCharsComms {
			if helpers.StringSliceAllUnique(checkChars) {
				return idx + 1
			}
		}
	}
	return -1
}

func findStartOfMessage(input string) int {
	for idx := range input {
		char := input[idx]
		processNewChar(char)
		if getNonEmptyLength(processed) == consecutiveCharsMessage {
			if helpers.StringSliceAllUnique(processed[0:consecutiveCharsMessage]) {
				return idx + 1
			}
		}
	}
	return -1
}

func part1(filename string) int {
	data, err := helpers.OpenAndReadFile(filename)
	if err != nil {
		panic(err)
	}
	strData := string(data)
	return findStartOfCommunication(strData)
}

func part2(filename string) int {
	data, err := helpers.OpenAndReadFile(filename)
	if err != nil {
		panic(err)
	}
	strData := string(data)
	return findStartOfMessage(strData)
}

func main() {
	p1 := part1("input.txt")
	fmt.Printf("Part 1: %d\n", p1)
	p2 := part2("input.txt")
	fmt.Printf("Part 2: %d\n", p2)
}
