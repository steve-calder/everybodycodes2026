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

	questList := [...]quests.Quest{
		&quests.Quest1{},
		&quests.Quest2{},
		&quests.Quest3{},
	}

	if len(questList) > 0 && len(os.Args) > 1 {
		i, err := strconv.ParseInt(os.Args[1], 10, 64)

		if err != nil {
			panic(err)
		}

		quests.ExecuteQuest(questList[i-1], notesQuery)
	} else {
		for _, quest := range questList {
			quests.ExecuteQuest(quest, notesQuery)
		}
	}
}
