package quests

import (
	"maps"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Quest4 struct{}

func (q *Quest4) getQuestId() int {
	return 4
}

func (q *Quest4) part1(input string) string {
	lines := strings.Split(input, "\n")
	regex := regexp.MustCompile(`id=(\d+) quest=(\d+) part=(\d+) time=(\d+)`)

	places := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)
		idx, _ := strconv.Atoi(matches[1])
		score, _ := strconv.Atoi(matches[4])
		places[idx-1] += score
	}

	slices.Sort(places)

	checksum := 0

	for i := 0; i < len(places); i++ {
		checksum += places[i] * (i + 1)
	}

	return strconv.Itoa(checksum)
}

type q4p2Participant struct {
	id     int
	events int
	score  int
}

func (q *Quest4) part2(input string) string {
	lines := strings.Split(input, "\n")
	regex := regexp.MustCompile(`id=(\d+) quest=(\d+) part=(\d+) time=(\d+)`)

	participants := map[int]*q4p2Participant{}
	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)
		id, _ := strconv.Atoi(matches[1])
		score, _ := strconv.Atoi(matches[4])

		if _, exists := participants[id]; !exists {
			participants[id] = &q4p2Participant{id: id, events: 0, score: 0}
		}
		participants[id].events++
		participants[id].score += score
	}

	participantIds := slices.Collect(maps.Keys(participants))
	slices.SortFunc(participantIds, func(a, b int) int {
		if participants[a].events > participants[b].events {
			return -1
		}

		if participants[a].events < participants[b].events {
			return 1
		}

		if participants[a].score < participants[b].score {
			return -1
		}

		if participants[a].score > participants[b].score {
			return 1
		}

		return 0
	})

	checksum := 0

	for i := 0; i < len(participantIds); i++ {
		checksum += participants[participantIds[i]].score * (i + 1)
	}

	return strconv.Itoa(checksum)
}

type q4p3Event struct {
	id    string
	part1 map[int]int
	part2 map[int]int
	part3 map[int]int
}

func (q *q4p3Event) Scores() *map[int]int {
	totalScores := map[int]int{}

	participantIds := slices.Collect(maps.Keys(q.part1))
	slices.SortFunc(participantIds, func(a, b int) int {
		if q.part1[a] < q.part1[b] {
			return -1
		}

		if q.part1[a] > q.part1[b] {
			return 1
		}

		return 0
	})

	decrement := 1
	previousTime := -1
	score := 51
	i := 0

	for score > 0 && i < len(participantIds) {
		participantTime := q.part1[participantIds[i]]
		if participantTime != previousTime {
			previousTime = participantTime
			score -= decrement
			if score < 0 {
				score = 0
			}
			decrement = 1

			totalScores[participantIds[i]] += score
		} else {
			decrement++
			totalScores[participantIds[i]] += score
		}

		i++
	}

	participantIds = slices.Collect(maps.Keys(q.part2))
	slices.SortFunc(participantIds, func(a, b int) int {
		aP1 := 0
		bP1 := 0

		if _, exists := q.part1[a]; exists {
			aP1 = q.part1[a]
		}

		if _, exists := q.part1[b]; exists {
			bP1 = q.part1[b]
		}

		if q.part2[a]+aP1 < q.part2[b]+bP1 {
			return -1
		}

		if q.part2[a]+aP1 > q.part2[b]+bP1 {
			return 1
		}

		return 0
	})

	decrement = 1
	previousTime = -1
	score = 101
	i = 0

	for score > 0 && i < len(participantIds) {
		p1ParticipantTime := 0
		if _, exists := q.part1[participantIds[i]]; exists {
			p1ParticipantTime = q.part1[participantIds[i]]
		}
		participantTime := q.part2[participantIds[i]] + p1ParticipantTime
		if participantTime != previousTime {
			previousTime = participantTime
			score -= decrement
			if score < 0 {
				score = 0
			}
			decrement = 1

			totalScores[participantIds[i]] += score
		} else {
			decrement++
			totalScores[participantIds[i]] += score
		}

		i++
	}

	participantIds = slices.Collect(maps.Keys(q.part3))
	slices.SortFunc(participantIds, func(a, b int) int {
		aP1 := 0
		aP2 := 0
		bP1 := 0
		bP2 := 0

		if _, exists := q.part1[a]; exists {
			aP1 = q.part1[a]
		}

		if _, exists := q.part1[b]; exists {
			bP1 = q.part1[b]
		}

		if _, exists := q.part2[a]; exists {
			aP2 = q.part2[a]
		}

		if _, exists := q.part2[b]; exists {
			bP2 = q.part2[b]
		}

		if q.part3[a]+aP1+aP2 < q.part3[b]+bP1+bP2 {
			return -1
		}

		if q.part3[a]+aP1+aP2 > q.part3[b]+bP1+bP2 {
			return 1
		}

		return 0
	})

	decrement = 1
	previousTime = -1
	score = 151
	i = 0

	for score > 0 && i < len(participantIds) {
		p1ParticipantTime := 0
		p2ParticipantTime := 0
		if _, exists := q.part1[participantIds[i]]; exists {
			p1ParticipantTime = q.part1[participantIds[i]]
		}
		if _, exists := q.part2[participantIds[i]]; exists {
			p2ParticipantTime = q.part2[participantIds[i]]
		}
		participantTime := q.part3[participantIds[i]] + p1ParticipantTime + p2ParticipantTime
		if participantTime != previousTime {
			previousTime = participantTime
			score -= decrement
			if score < 0 {
				score = 0
			}
			decrement = 1

			totalScores[participantIds[i]] += score
		} else {
			decrement++
			totalScores[participantIds[i]] += score
		}

		i++
	}

	return &totalScores
}

func (q *Quest4) part3(input string) string {
	lines := strings.Split(input, "\n")
	regex := regexp.MustCompile(`id=(\d+) quest=(\d+) part=(\d+) time=(\d+)`)

	participantScores := map[int]int{}
	events := map[string]*q4p3Event{}

	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)
		id, _ := strconv.Atoi(matches[1])
		quest := matches[2]
		part := matches[3]
		score, _ := strconv.Atoi(matches[4])

		if _, exists := events[quest]; !exists {
			events[quest] = &q4p3Event{id: quest, part1: map[int]int{}, part2: map[int]int{}, part3: map[int]int{}}
		}

		switch part {
		case "1":
			events[quest].part1[id] = score
		case "2":
			events[quest].part2[id] = score
		case "3":
			events[quest].part3[id] = score
		}

		if _, exists := participantScores[id]; !exists {
			participantScores[id] = 0
		}
	}

	for _, event := range events {
		results := event.Scores()
		for id, score := range *results {
			participantScores[id] += score
		}
	}

	participantIds := slices.Collect(maps.Keys(participantScores))
	slices.SortFunc(participantIds, func(a, b int) int {
		if participantScores[a] > participantScores[b] {
			return -1
		}

		if participantScores[a] < participantScores[b] {
			return 1
		}

		return 0
	})

	checksum := 0
	previousScore := -1
	decrement := 0

	for i := 0; i < len(participantIds); i++ {
		if participantScores[participantIds[i]] == previousScore {
			decrement++
		} else {
			decrement = 0
		}

		previousScore = participantScores[participantIds[i]]
		checksum += participantScores[participantIds[i]] * (i + 1 - decrement)
	}

	return strconv.Itoa(checksum)
}
