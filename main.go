package main

import (
	"os"
	"strconv"

	"github.com/steve-calder/everybodycodes2026/notes"
	"github.com/steve-calder/everybodycodes2026/quests"
	"gopkg.in/yaml.v3"
)

var config notes.Config

func loadConfig() error {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, &config)
}

func main() {
	if err := loadConfig(); err != nil {
		panic(err)
	}

	notesQuery := notes.InitNotesQuery(&config)

	quests := [...]quests.Quests{
		&quests.Quest1{NotesQuery: notesQuery},
	}

	if len(quests) > 0 && len(os.Args) > 1 {
		i, err := strconv.ParseInt(os.Args[1], 10, 64)

		if err != nil {
			panic(err)
		}

		quests[i-1].Execute()
	} else {
		for _, quest := range quests {
			quest.Execute()
		}
	}
}
