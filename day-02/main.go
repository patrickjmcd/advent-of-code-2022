package main

import (
	"fmt"
	"github.com/patrickjmcd/advent-of-code-2022/helpers"
	"strings"
)

// points per selection
var (
	rock     = 1
	paper    = 2
	scissors = 3
)

// points per outcome
var (
	lose = 0
	draw = 3
	win  = 6
)

func getNameForCode(code string) string {
	switch code {
	case "A", "X":
		return "rock"
	case "B", "Y":
		return "paper"
	case "C", "Z":
		return "scissors"
	}
	return ""
}

func getCodeForName(name string) string {
	switch name {
	case "rock":
		return "A"
	case "paper":
		return "B"
	case "scissors":
		return "C"
	}
	return ""
}

func rockPaperScissorsGame(them, me string) int {
	theirPick := getNameForCode(them)
	myPick := getNameForCode(me)

	switch theirPick {
	case "rock":
		switch myPick {
		case "rock":
			return rock + draw
		case "paper":
			return paper + win
		case "scissors":
			return scissors + lose
		}
	case "paper":
		switch myPick {
		case "rock":
			return rock + lose
		case "paper":
			return paper + draw
		case "scissors":
			return scissors + win
		}
	case "scissors":
		switch myPick {
		case "rock":
			return rock + win
		case "paper":
			return paper + lose
		case "scissors":
			return scissors + draw
		}
	}
	return 0
}

func getOutcomeForCode(code string) string {
	switch code {
	case "X":
		return "lose"
	case "Y":
		return "draw"
	case "Z":
		return "win"
	}
	return ""
}

func getYourSelection(opponent, outcome string) string {
	switch outcome {
	case "win":
		switch opponent {
		case "rock":
			return "paper"
		case "paper":
			return "scissors"
		case "scissors":
			return "rock"
		}
	case "lose":
		switch opponent {
		case "rock":
			return "scissors"
		case "paper":
			return "rock"
		case "scissors":
			return "paper"
		}
	case "draw":
		return opponent
	}
	return ""
}

func part1(filename string) int {
	data, err := helpers.OpenAndReadFile(filename)
	if err != nil {
		panic(err)
	}

	strInput := string(data)
	individualGames := strings.Split(strInput, "\n")

	total := 0
	for i := range individualGames {
		if len(individualGames[i]) > 0 {
			values := strings.Split(individualGames[i], " ")
			total += rockPaperScissorsGame(values[0], values[1])
		}
	}
	fmt.Printf("Total points part 1 is %d\n", total)
	return total
}

func part2(filename string) int {
	data, err := helpers.OpenAndReadFile(filename)
	if err != nil {
		panic(err)
	}
	strInput := string(data)
	individualGames := strings.Split(strInput, "\n")

	total := 0
	for i := range individualGames {
		if len(individualGames[i]) > 0 {
			values := strings.Split(individualGames[i], " ")
			opponentChoice := getNameForCode(values[0])
			outcome := getOutcomeForCode(values[1])
			yourChoice := getYourSelection(opponentChoice, outcome)
			yourCode := getCodeForName(yourChoice)
			total += rockPaperScissorsGame(values[0], yourCode)
		}
	}
	fmt.Printf("Total points part 2 is %d\n", total)
	return total
}

func main() {
	part1("./input.txt")
	part2("./input.txt")
}
