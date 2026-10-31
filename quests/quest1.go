package quests

import (
	"strconv"
	"strings"
)

type Quest1 struct{}

func (q *Quest1) getQuestId() int {
	return 1
}

func (q *Quest1) part1(input string) string {
	nums := parseInput(input)

	total := 0
	square := 1

	for _, num := range nums {
		for square * square < num {
			square++
		}
		total += square * square - num
	}

	return strconv.Itoa(total)
}

func (q *Quest1) part2(input string) string {
	nums := parseInput(input)

	total := 0
	square := 3
	for _, num := range nums {
		for square * 4 - 4 < num {
			square++
		}
		total += square * 4 - 4 - num
	}
	return strconv.Itoa(total)
}

func (q *Quest1) part3(input string) string {
	nums := parseInput(input)

	total := 0
	square := 3
	for _, num := range nums {
		for square * square - 1 < num {
			square += 2
		}
		total += square * square - 1 - num
	}
	return strconv.Itoa(total)
}

func parseInput(input string) []int {
	numTokens := strings.Split(input, ",")
	nums := make([]int, len(numTokens))
	for i, token := range numTokens {
		nums[i], _ = strconv.Atoi(token)
	}
	return nums
}