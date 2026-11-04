package quests

import (
	"maps"
	"regexp"
	"strconv"
	"strings"
)

type Quest7 struct{}

func (q *Quest7) getQuestId() int {
	return 7
}

func (q *Quest7) part1(input string) string {
	lines := strings.Split(input, "\n")
	regex := regexp.MustCompile(`id=(\d+), energy=(\d+), connections=(.*)`)

	lightning := map[int][]int{}
	energy := map[int]int{}

	largestId := 0
	largestConnections := 0

	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)
		idx, _ := strconv.Atoi(matches[1])
		energyVal, _ := strconv.Atoi(matches[2])
		connections := strings.Split(matches[3], ",")

		connectionList := []int{}

		for _, conn := range connections {
			connIdx, _ := strconv.Atoi(conn)
			connectionList = append(connectionList, connIdx)
		}

		energy[idx] = energyVal

		if len(connectionList) > largestConnections {
			largestConnections = len(connectionList)
			largestId = idx
		}

		lightning[idx] = connectionList
	}

	stack := []int{largestId}
	impulse := 5

	for impulse > 0 {
		nextStack := []int{}

		for _, current := range stack {
			for _, conn := range lightning[current] {
				nextStack = append(nextStack, conn)
			}

			energy[current] += impulse
		}

		stack = nextStack
		impulse--
	}

	totalEnergy := 0

	for _, e := range energy {
		totalEnergy += e
	}

	return strconv.Itoa(totalEnergy)
}

func (q *Quest7) part2(input string) string {
	lines := strings.Split(input, "\n")
	regex := regexp.MustCompile(`id=(\d+), energy=(\d+), connections=(.*)`)

	lightning := map[int64][]int64{}
	energy := map[int64]int64{}

	largestId := int64(0)
	largestConnections := int64(0)

	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)
		idx, _ := strconv.Atoi(matches[1])
		energyVal, _ := strconv.Atoi(matches[2])
		connections := strings.Split(matches[3], ",")

		connectionList := []int64{}

		for _, conn := range connections {
			connIdx, _ := strconv.Atoi(conn)
			connectionList = append(connectionList, int64(connIdx))
		}

		energy[int64(idx)] = int64(energyVal)

		if int64(len(connectionList)) > largestConnections {
			largestConnections = int64(len(connectionList))
			largestId = int64(idx)
		}

		lightning[int64(idx)] = connectionList
	}

	toVisitMap := map[int64]int64{int64(largestId): int64(1)}
	impulse := 10

	for impulse > 0 {
		newToVisitMap := map[int64]int64{}

		for i, v := range toVisitMap {
			energy[i] += int64(impulse) * v

			for _, conn := range lightning[i] {
				newToVisitMap[conn] += v
			}
		}

		toVisitMap = newToVisitMap
		impulse--
	}

	totalEnergy := int64(0)

	for _, e := range energy {
		totalEnergy += e
	}

	return strconv.FormatInt(totalEnergy, 10)
}

func (q *Quest7) part3(input string) string {
	lines := strings.Split(input, "\n")
	regex := regexp.MustCompile(`id=(\d+), energy=(\d+), connections=(.*)`)

	lightning := map[int64][]int64{}
	baseLineEnergy := map[int64]int64{}

	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)
		idx, _ := strconv.Atoi(matches[1])
		energyVal, _ := strconv.Atoi(matches[2])
		connections := strings.Split(matches[3], ",")

		connectionList := []int64{}

		for _, conn := range connections {
			connIdx, _ := strconv.Atoi(conn)
			connectionList = append(connectionList, int64(connIdx))
		}

		baseLineEnergy[int64(idx)] = int64(energyVal)

		lightning[int64(idx)] = connectionList
	}

	maxKey := int64(0)
	largestEnergy := int64(0)

	for largestId := 1; largestId <= 100; largestId++ {
		energy := maps.Clone(baseLineEnergy)

		toVisitMap := map[int64]int64{int64(largestId): int64(1)}

		impulseList := []int64{10}
		toVisitMapList := []map[int64]int64{toVisitMap}
		exploded := map[int64]bool{}

		totalCharge := 10

		for totalCharge > 0 {
			for y := 0; y < len(toVisitMapList); y++ {
				if impulseList[y] <= 0 {
					continue
				}

				newToVisitMap := map[int64]int64{}

				for i, v := range toVisitMapList[y] {
					if _, exists := exploded[i]; exists {
						continue
					}

					energy[i] += int64(impulseList[y]) * v

					for _, conn := range lightning[i] {
						if _, exists := exploded[conn]; !exists {
							newToVisitMap[conn] += v
						}
					}
				}

				impulseList[y]--
				totalCharge--
				toVisitMapList[y] = newToVisitMap
			}

			for j := range energy {
				if energy[j] > 1000000 {
					if _, exists := exploded[j]; !exists {
						exploded[j] = true
						impulseList = append(impulseList, 10)

						explodedNewToVisitMap := map[int64]int64{}
						for _, conn := range lightning[j] {
							if _, exists := exploded[conn]; !exists {
								explodedNewToVisitMap[conn] = 1
							}
						}

						toVisitMapList = append(toVisitMapList, explodedNewToVisitMap)
						totalCharge += 10
					}
				}
			}
		}

		totalEnergy := int64(0)

		for k, e := range energy {
			if _, exists := exploded[k]; !exists {
				totalEnergy += e
			}
		}

		if totalEnergy > largestEnergy {
			largestEnergy = totalEnergy
			maxKey = int64(largestId)
		}
	}

	println("Max key:", maxKey, "Largest energy:", largestEnergy)

	return strconv.FormatInt(largestEnergy, 10)
}
