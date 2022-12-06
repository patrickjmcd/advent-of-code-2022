package main

import (
	"fmt"
	"github.com/patrickjmcd/advent-of-code-2022/helpers"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type command struct {
	quantity int
	from     int
	to       int
}

type container struct {
	label  string
	stack  int
	height int
}

type byHeight []container

func (c byHeight) Len() int {
	return len(c)
}
func (c byHeight) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c byHeight) Less(i, j int) bool {
	return c[i].height < c[j].height
}

func parseCommand(line string) command {
	var c command
	splitLine := strings.Split(line, " ")
	c.quantity, _ = strconv.Atoi(splitLine[1])
	c.from, _ = strconv.Atoi(splitLine[3])
	c.to, _ = strconv.Atoi(splitLine[5])
	return c
}

func parseDrawing(drawing string) ([]container, int) {
	lines := strings.Split(drawing, "\n")
	linesWithoutLastLine := lines[:len(lines)-1]
	maxHeight := len(linesWithoutLastLine)
	lastLine := lines[len(lines)-1]
	r := regexp.MustCompile(`\d+`)
	matches := r.FindAllString(lastLine, -1)
	numStacks := len(matches)
	fmt.Printf("Drawing has %d stacks\n", numStacks)
	var containers []container

	for l, line := range linesWithoutLastLine {
		if len(line) > 0 {
			startIdx := 1
			for i := 0; i < numStacks; i++ {
				selector := startIdx + 4*i
				if selector < len(line) {
					value := line[selector]
					if value != ' ' {
						c := container{label: string(value), stack: i + 1, height: maxHeight - l}
						containers = append(containers, c)
					}
				}
			}
		}
	}
	return containers, numStacks
}

func makeStacksFromContainers(containers []container) map[int][]string {

	stacks := map[int][]container{}
	for _, c := range containers {
		if _, ok := stacks[c.stack]; !ok {
			stacks[c.stack] = []container{}
		}
		stacks[c.stack] = append(stacks[c.stack], c)
	}

	sortedStacks := map[int][]string{}

	for _, stack := range stacks {
		sort.Sort(sort.Reverse(byHeight(stack)))
		if _, ok := sortedStacks[stack[0].stack]; !ok {
			sortedStacks[stack[0].stack] = []string{}
		}
		for _, c := range stack {
			sortedStacks[stack[0].stack] = append(sortedStacks[stack[0].stack], c.label)
		}
	}

	return sortedStacks
}

func copyStr(a []string) []string {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = v
	}
	return b
}

func reverse(s []string) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[len(s)-1-i] = v
	}
	return r
}

func moveFromTo(from, to, quantity int, stacks *map[int][]string, withReversing bool) {
	var newStack []string
	stk := *stacks

	fromStack := copyStr(stk[from])
	toStack := copyStr(stk[to])

	if len(fromStack) >= quantity {
		moved := copyStr(fromStack[:quantity])
		if withReversing {
			moved = reverse(moved)
		}
		remaining := copyStr(fromStack[quantity:])

		newStack = append(moved, toStack...)
		stk[from] = remaining
		stk[to] = newStack
	}
	stacks = &stk
}

func part1(filename string, withReversing bool) (string, map[int][]string) {
	data, err := helpers.OpenAndReadFile(filename)
	if err != nil {
		panic(err)
	}
	strData := string(data)

	splitData := strings.Split(strData, "\n\n")
	drawing := splitData[0]
	instructions := splitData[1]

	containers, numStacks := parseDrawing(drawing)
	stacks := makeStacksFromContainers(containers)

	for _, line := range strings.Split(instructions, "\n") {
		if len(line) > 0 {
			cmd := parseCommand(line)
			moveFromTo(cmd.from, cmd.to, cmd.quantity, &stacks, withReversing)
		}
	}
	topLine := ""
	for i := 1; i <= numStacks; i++ {
		topLine += stacks[i][0]
	}
	fmt.Printf("Top line: %s\n", topLine)
	return topLine, stacks
}

func main() {
	fmt.Printf("Part1: ")
	part1("input.txt", true)
	fmt.Printf("Part2: ")
	part1("input.txt", false)
}
