package quests

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/steve-calder/everybodycodes2026/notes"
)

type Quest1 struct{
	NotesQuery *notes.NotesQuery
}

func (q *Quest1) Execute() {
	notes, err := q.NotesQuery.GetInput(1)
	if err != nil {
		fmt.Printf("Error retrieving notes for Quest 1: %v\n", err)
		return
	}

	fmt.Printf("Quest 1, Part 1: %d\n", q.part1(notes.Part1))
	fmt.Printf("Quest 1, Part 2: %d\n", q.part2(notes.Part2))
	fmt.Printf("Quest 1, Part 3: %d\n", q.part3(notes.Part3))
}

func (q *Quest1) part1(input string) int {
	nums := parseInput(input)

	total := 0
	square := 1

	for _, num := range nums {
		for square * square < num {
			square++
		}
		total += square * square - num
	}

	return total
}

func (q *Quest1) part2(input string) int {
	nums := parseInput(input)

	total := 0
	square := 3
	for _, num := range nums {
		for square * 4 - 4 < num {
			square++
		}
		total += square * 4 - 4 - num
	}
	return total
}

func (q *Quest1) part3(input string) int {
	nums := parseInput(input)

	total := 0
	square := 3
	for _, num := range nums {
		for square * square - 1 < num {
			square += 2
		}
		total += square * square - 1 - num
	}
	return total
}

func parseInput(input string) []int {
	numTokens := strings.Split(input, ",")
	nums := make([]int, len(numTokens))
	for i, token := range numTokens {
		nums[i], _ = strconv.Atoi(token)
	}
	return nums
}