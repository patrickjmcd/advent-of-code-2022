package main

import (
	"fmt"
	"github.com/patrickjmcd/advent-of-code-2022/helpers"
	"strings"
)

var priority = "0abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getLetterPriority(letter string) int {
	return strings.Index(priority, letter)
}

func splitStringInHalf(str string) (string, string) {
	mid := len(str) / 2
	return str[:mid], str[mid:]
}

func findCommonLetters(str1, str2 string) []string {
	var commonLetters []string
	for i := range str1 {
		if strings.Contains(str2, string(str1[i])) {
			if !helpers.StringInSlice(string(str1[i]), commonLetters) {
				commonLetters = append(commonLetters, string(str1[i]))
			}
		}
	}
	return commonLetters
}

func findCommonLettersThree(str1, str2, str3 string) []string {
	var commonLetters []string
	for i := range str1 {
		if strings.Contains(str2, string(str1[i])) && strings.Contains(str3, string(str1[i])) {
			if !helpers.StringInSlice(string(str1[i]), commonLetters) {
				commonLetters = append(commonLetters, string(str1[i]))
			}
		}
	}
	return commonLetters
}

func sumOfLetterPriority(str []string) int {
	sum := 0
	for i := range str {
		sum += getLetterPriority(str[i])
	}
	return sum
}

func part1(filename string) int {
	data, err := helpers.OpenAndReadFile(filename)
	if err != nil {
		panic(err)
	}
	strData := string(data)

	//split the file by lines
	lines := strings.Split(strData, "\n")

	total := 0
	// split each line in half
	for i := range lines {
		if len(lines[i]) > 0 {
			firstHalf, secondHalf := splitStringInHalf(lines[i])
			// find the common letters
			commonLetters := findCommonLetters(firstHalf, secondHalf)
			// sum the priority of the common letters
			sum := sumOfLetterPriority(commonLetters)
			total += sum
		}
	}
	fmt.Printf("Total points part 1 is %d\n", total)
	return total
}

func part2(filename string) int {
	// read the input file
	data, err := helpers.OpenAndReadFile(filename)
	if err != nil {
		panic(err)
	}
	strData := string(data)

	// split the file by lines
	lines := strings.Split(strData, "\n")

	//group the lines in groups of 3
	var groups [][]string
	for i := 0; i < len(lines); i += 3 {
		if len(lines[i]) > 0 {
			groups = append(groups, lines[i:i+3])
		}
	}

	total := 0

	for i := range groups {
		// find the common letters between each line in the group
		commonLetters := findCommonLettersThree(groups[i][0], groups[i][1], groups[i][2])
		// sum the priority of the common letters
		sum := sumOfLetterPriority(commonLetters)
		total += sum
	}
	fmt.Printf("Total points part 2 is %d\n", total)
	return total
}

func main() {
	part1("input.txt")
	part2("input.txt")
}
