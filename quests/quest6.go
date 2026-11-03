package quests

import (
	"strconv"
	"strings"
)

type Quest6 struct{}

func (q *Quest6) getQuestId() int {
	return 6
}

type ingredient struct {
	name        string
	ingredients *map[string]*ingredient
}

func (q *Quest6) part1(input string) string {
	lines := strings.Split(input, "\n")

	ingredients := make(map[string]*ingredient)

	for _, line := range lines {
		mainIngredient := strings.Split(line, ":")[0]
		subIngredients := strings.Split(strings.Split(line, ":")[1], ",")

		subIngredientMap := make(map[string]*ingredient)

		for _, subIngredient := range subIngredients {
			if _, ok := ingredients[subIngredient]; !ok {
				freshSubIngredientMap := make(map[string]*ingredient)
				subIngredientMap[subIngredient] = &ingredient{name: subIngredient, ingredients: &freshSubIngredientMap}
				ingredients[subIngredient] = subIngredientMap[subIngredient]
			} else {
				subIngredientMap[subIngredient] = ingredients[subIngredient]
			}
		}

		if _, ok := ingredients[mainIngredient]; !ok {
			ingredients[mainIngredient] = &ingredient{name: mainIngredient, ingredients: &subIngredientMap}
		} else {
			ingredients[mainIngredient].ingredients = &subIngredientMap
		}
	}

	uniqueIngredients := make(map[string]bool)

	var countIngredients func(ing *ingredient)
	countIngredients = func(ing *ingredient) {
		if _, ok := uniqueIngredients[ing.name]; ok {
			return
		}

		uniqueIngredients[ing.name] = true
		for _, subIng := range *ing.ingredients {
			countIngredients(subIng)
		}
	}

	countIngredients(ingredients["NULL"])

	total := len(uniqueIngredients)
	total--

	return strconv.Itoa(total)
}

func (q *Quest6) part2(input string) string {
	lines := strings.Split(input, "\n")

	ingredients := make(map[string][]string)

	for _, line := range lines {
		mainIngredient := strings.Split(line, ":")[0]
		subIngredients := strings.Split(strings.Split(line, ":")[1], ",")

		for _, subIngredient := range subIngredients {
			if _, ok := ingredients[subIngredient]; !ok {
				list := []string{mainIngredient}
				ingredients[subIngredient] = list
			} else {
				ingredients[subIngredient] = append(ingredients[subIngredient], mainIngredient)
			}
		}
	}

	uniqueIngredients := make(map[string]bool)
	stack := []string{"NULL"}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if _, ok := uniqueIngredients[current]; ok {
			continue
		}

		uniqueIngredients[current] = true

		for _, parent := range ingredients[current] {
			stack = append(stack, parent)
		}
	}

	total := len(uniqueIngredients)
	total--

	return strconv.Itoa(total)
}

func (q *Quest6) part3(input string) string {
	lines := strings.Split(input, "\n")

	ingredients := make(map[string][]string)

	for _, line := range lines {
		mainIngredient := strings.Split(line, ":")[0]
		subIngredients := strings.Split(strings.Split(line, ":")[1], ",")

		for _, subIngredient := range subIngredients {
			if _, ok := ingredients[subIngredient]; !ok {
				list := []string{mainIngredient}
				ingredients[subIngredient] = list
			} else {
				ingredients[subIngredient] = append(ingredients[subIngredient], mainIngredient)
			}
		}
	}

	nullCounts := make(map[string]int)

	nulls := []string{"NULLv1", "NULLv2", "NULLv3"}

	for _, null := range nulls {
		uniqueIngredients := make(map[string]bool)
		stack := []string{null}

		for len(stack) > 0 {
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if _, ok := uniqueIngredients[current]; ok {
				continue
			}

			uniqueIngredients[current] = true
			nullCounts[current]++

			for _, parent := range ingredients[current] {
				stack = append(stack, parent)
			}
		}
	}

	total := 0
	for _, count := range nullCounts {
		if count > 1 {
			total++
		}
	}

	return strconv.Itoa(total)
}
