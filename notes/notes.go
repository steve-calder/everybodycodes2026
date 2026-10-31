package notes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Notes struct {
	Part1 string `json:"1"`
	Part2 string `json:"2"`
	Part3 string `json:"3"`
}

type NotesQuery struct {
	config *Config
}

func InitNotesQuery(config *Config) *NotesQuery {
	return &NotesQuery{
		config: config,
	}
}

func (nq *NotesQuery) GetInput(quest int) (*Notes, error) {
	cachePath := filepath.Join("cache", fmt.Sprintf("%d-%d-%d.json", nq.config.Event, quest, nq.config.Seed))
	data, err := os.ReadFile(cachePath)
	if err == nil {
		var notes Notes
		if err := json.Unmarshal(data, &notes); err != nil {
			return nil, fmt.Errorf("decode cached notes: %w", err)
		}
		return &notes, nil
	}
	if !os.IsNotExist(err) {
		return nil, fmt.Errorf("read cached notes: %w", err)
	}

	url := fmt.Sprintf("https://everybody.codes/assets/%d/%d/input/%d.json", nq.config.Event, quest, nq.config.Seed)
	client := &http.Client{Timeout: 30 * time.Second}
	response, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("retrieve notes: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("retrieve notes: HTTP %d", response.StatusCode)
	}

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read notes: %w", err)
	}
	var notes Notes
	if err := json.Unmarshal(data, &notes); err != nil {
		return nil, fmt.Errorf("decode notes: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(cachePath), 0700); err != nil {
		return nil, fmt.Errorf("create cache directory: %w", err)
	}
	if err := os.WriteFile(cachePath, data, 0600); err != nil {
		return nil, fmt.Errorf("cache notes: %w", err)
	}
	return &notes, nil
}
