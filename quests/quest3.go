package quests

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/steve-calder/everybodycodes2026/utils"
)

type Quest3 struct{}

func (q *Quest3) getQuestId() int {
	return 3
}

func (q *Quest3) part1(input string) string {
	lines := strings.Split(input, "\n")

	total := 0.0

	for _, line := range lines {
		score := 0.0

		power := 3
		for i := range 4 {
			weight := int(line[4+i] - '0')
			digit := int(line[9+i] - '0')
			score += math.Pow(float64(weight), float64(power)) * float64(digit)
			power--
		}

		total += score
	}

	return fmt.Sprintf("%.0f", total)
}

func (q *Quest3) part2(input string) string {
	lines := strings.Split(input, "\n")

	total := 0.0

	for _, line := range lines {
		score := 0.0

		weights, digits := parseLinePart2(line)
		power := 0

		weights = utils.ReverseString(weights)
		digits = utils.ReverseString(digits)

		weightQueue := make(chan int, len(weights))
		for _, weight := range weights {
			weightQueue <- int(weight - '0')
		}

		for i := range len(digits) {
			weight := <-weightQueue
			digit := int(digits[i] - '0')
			score += math.Pow(float64(weight), float64(power)) * float64(digit)
			power++
			weightQueue <- weight
		}

		total += score
	}

	return fmt.Sprintf("%.0f", total)
}

func parseLinePart2(line string) (string, string) {
	regex := regexp.MustCompile(`ISO-(\d+) (\d+)`)
	matches := regex.FindStringSubmatch(line)
	return matches[1], matches[2]
}

func (q *Quest3) part3(input string) string {
	lines := strings.Split(input, "\n")

	total := 0.0

	for _, line := range lines {
		weights, digits := parseLinePart3(line)
		weights = utils.ReverseString(weights) // reversing the weights makes it easier to get the right index
		digitsNum, _ := strconv.Atoi(digits)
		cache := make(map[string]float64)

		var calc func(string) float64
		calc = func(d string) float64 {
			s := 0.0

			if len(d) > 1 {
				if val, exists := cache[d[1:]]; exists {
					s += val
				} else {
					s += calc(d[1:])
				}
			}

			power := len(d) - 1

			weightIdx := (len(d) - 1) % len(weights)

			weight := int(weights[weightIdx] - '0')
			digit := int(d[0] - '0')
			s += math.Pow(float64(weight), float64(power)) * float64(digit)

			cache[d] = s
			return s
		}

		maxScore := calc(digits)
		digitsIsBiggest := true

		for i := maxScore; i < float64(digitsNum); i++ {
			score := calc(strconv.Itoa(int(i)))

			if score == maxScore {
				digitsIsBiggest = false
				break
			}
		}

		if digitsIsBiggest {
			total += maxScore
		}
	}

	return fmt.Sprintf("%.0f", total)
}

func parseLinePart3(line string) (string, string) {
	regex := regexp.MustCompile(`ISO2-(\d+) (\d+)`)
	matches := regex.FindStringSubmatch(line)
	return matches[1], matches[2]
}
