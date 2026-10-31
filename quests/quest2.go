package quests

import (
	"strconv"
	"strings"
)

type Quest2 struct{}

func (q *Quest2) getQuestId() int {
	return 2
}

/* Just manually swapped all the places */
func (q *Quest2) part1(input string) string {
	lines := strings.Split(input, "\n")
	total := 0

	for _, line := range lines {
		ids := map[rune]int{'1': 1, '2': 2, '3': 3, '4': 4, '5': 5}

		rules := strings.Split(line, " ")

		for _, rule := range rules {
			id1 := rune(rule[0])
			id2 := rune(rule[2])
			operator := rune(rule[1])

			if operator == '<' {
				if ids[id1] > ids[id2] {
					id1Idx := ids[id1]
					id2Idx := ids[id2]

					for id, idx := range ids {
						if idx < id1Idx && idx >= id2Idx {
							ids[id] = idx + 1
						}
					}

					ids[id1] = id2Idx
				}
			} else {
				if ids[id2] > ids[id1] {
					id1Idx := ids[id1]
					id2Idx := ids[id2]

					for id, idx := range ids {
						if idx < id2Idx && idx >= id1Idx {
							ids[id] = idx + 1
						}
					}

					ids[id2] = id1Idx
				}
			}
		}

		runes := []rune(".....")
		for id, idx := range ids {
			runes[idx-1] = id
		}

		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		finalString := string(runes)
		final, _ := strconv.Atoi(finalString)
		total += final
	}
	return strconv.Itoa(total)
}

/* Just manually swapped all the places but re-ran it multiple times to ensure the rules were applied correctly */
func (q *Quest2) part2(input string) string {
	lines := strings.Split(input, "\n")
	total := 0

	for _, line := range lines {
		ids := map[rune]int{'1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}

		rules := strings.Split(line, " ")

		swapped := true

		for swapped {
			swapped = false
			for _, rule := range rules {
				id1 := rune(rule[0])
				id2 := rune(rule[2])
				operator := rune(rule[1])

				if operator == '<' {
					if ids[id1] > ids[id2] {
						swapped = true
						id1Idx := ids[id1]
						id2Idx := ids[id2]

						for id, idx := range ids {
							if idx < id1Idx && idx >= id2Idx {
								ids[id] = idx + 1
							}
						}

						ids[id1] = id2Idx
					}
				} else {
					if ids[id2] > ids[id1] {
						swapped = true
						id1Idx := ids[id1]
						id2Idx := ids[id2]

						for id, idx := range ids {
							if idx < id2Idx && idx >= id1Idx {
								ids[id] = idx + 1
							}
						}

						ids[id2] = id1Idx
					}
				}
			}
		}

		runes := []rune(".........")
		for id, idx := range ids {
			runes[idx-1] = id
		}

		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		finalString := string(runes)
		final, _ := strconv.Atoi(finalString)
		total += final
	}
	return strconv.Itoa(total)
}

/* The approach for part2 still works but it's slower (6.1ms vs < 1ms).  Research shows I can do Kahn's algorithm, so I tried to implement that which came in at 3.3ms */
func (q *Quest2) part3(input string) string {
	lines := strings.Split(input, "\n")
	total := 0

	for _, line := range lines {
		// build the adjacency list.  basically all lower digits are pointing to higher digits
		adj := [9][]int{}
		for i := range adj {
			adj[i] = []int{}
		}

		rules := strings.Split(line, " ")

		for _, rule := range rules {
			id1, _ := strconv.Atoi(string(rule[0]))
			id2, _ := strconv.Atoi(string(rule[2]))
			operator := rune(rule[1])

			if operator == '<' {
				adj[id1-1] = append(adj[id1-1], id2-1)
			} else {
				adj[id2-1] = append(adj[id2-1], id1-1)
			}
		}

		// count the number of incoming edges for each node
		indegree := [9]int{}
		for i := 0; i < 9; i++ {
			for _, j := range adj[i] {
				indegree[j]++
			}
		}

		// queue up the nodes with no incoming edges
		queue := make(chan int, 9)
		for i := 0; i < 9; i++ {
			if indegree[i] == 0 {
				queue <- i
			}
		}

		answer := []string{}

		// Kahn algorithm - https://www.geeksforgeeks.org/dsa/topological-sorting-indegree-based-solution/
		for len(queue) > 0 {
			current := <-queue

			answer = append(answer, strconv.Itoa(current+1))
			for _, neighbor := range adj[current] {
				indegree[neighbor]--
				if indegree[neighbor] == 0 {
					queue <- neighbor
				}
			}
		}

		for i, j := 0, len(answer)-1; i < j; i, j = i+1, j-1 {
			answer[i], answer[j] = answer[j], answer[i]
		}

		finalString := strings.Join(answer, "")
		final, _ := strconv.Atoi(finalString)
		total += final
	}

	return strconv.Itoa(total)
}
