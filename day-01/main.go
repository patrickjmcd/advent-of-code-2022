package main

import (
	"fmt"
	"github.com/patrickjmcd/advent-of-code-2022/helpers"
	"sort"
	"strconv"
	"strings"
)

type elf struct {
	id       int
	calories int
}

type byCalories []elf

func (a byCalories) Len() int           { return len(a) }
func (a byCalories) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byCalories) Less(i, j int) bool { return a[i].calories < a[j].calories }

func topElf() {
	data, err := helpers.OpenAndReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	var max elf

	strInput := string(data)
	individualElves := strings.Split(strInput, "\n\n")
	for i := range individualElves {
		values := strings.Split(individualElves[i], "\n")
		sum := 0
		for j := range values {
			if len(values[j]) > 0 {
				intVal, err := strconv.Atoi(values[j])
				if err != nil {
					panic(err)
				}
				sum += intVal
			}
		}
		if sum > max.calories {
			max.calories = sum
			max.id = i + 1
		}
	}

	fmt.Printf("Elf %d has the most calories with %d\n", max.id, max.calories)
}

func top3elves() {
	data, err := helpers.OpenAndReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	var totals []elf

	strInput := string(data)
	individualElves := strings.Split(strInput, "\n\n")
	for i := range individualElves {
		values := strings.Split(individualElves[i], "\n")
		sum := 0
		for j := range values {
			if len(values[j]) > 0 {
				intVal, err := strconv.Atoi(values[j])
				if err != nil {
					panic(err)
				}
				sum += intVal
			}
		}
		totals = append(totals, elf{id: i + 1, calories: sum})
	}
	sort.Sort(sort.Reverse(byCalories(totals)))
	fmt.Printf("Sum of the top 3 elves is %d\n", totals[0].calories+totals[1].calories+totals[2].calories)

}

func main() {
	topElf()
	top3elves()
}
