package quests

import (
	"fmt"

	"github.com/steve-calder/everybodycodes2026/notes"
)

type Quest interface {
	getQuestId() int
	part1(input string) string
	part2(input string) string
	part3(input string) string
}

func ExecuteQuest(quest Quest, notesQuery *notes.NotesQuery) {
	questId := quest.getQuestId()

	notes, err := notesQuery.GetInput(questId)
	if err != nil {
		fmt.Printf("Error retrieving notes for Quest %d: %v\n", questId, err)
		return
	}

	part1Input, err := notes.Decode(1)
	if err != nil {
		fmt.Printf("Error decoding Quest %d Part 1: %v\n", questId, err)
	} else {
		fmt.Printf("Quest %d, Part 1: %s\n", questId, quest.part1(part1Input))
	}
	
	part2Input, err := notes.Decode(2)
	if err != nil {
		fmt.Printf("Error decoding Quest %d Part 2: %v\n", questId, err)
	} else {
		fmt.Printf("Quest %d, Part 2: %s\n", questId, quest.part2(part2Input))
	}

	part3Input, err := notes.Decode(3)
	if err != nil {
		fmt.Printf("Error decoding Quest %d Part 3: %v\n", questId, err)
	} else {
		fmt.Printf("Quest %d, Part 3: %s\n", questId, quest.part3(part3Input))
	}
}