package notes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"unicode/utf8"
)

type Notes struct {
	Part1  string `json:"1"`
	Part2  string `json:"2"`
	Part3  string `json:"3"`
	config *Config
	quest  int
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
		notes := Notes{config: nq.config, quest: quest}
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
	notes := Notes{config: nq.config, quest: quest}
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

// Decode returns the decrypted input for part 1, 2, or 3.
func (n *Notes) Decode(part int) (string, error) {
	if part < 1 || part > 3 {
		return "", fmt.Errorf("part must be between 1 and 3")
	}
	if n.config == nil || n.quest < 1 {
		return "", fmt.Errorf("load notes with GetInput before decoding")
	}
	key, err := n.getKey(part)
	if err != nil {
		return "", err
	}
	input := []string{n.Part1, n.Part2, n.Part3}[part-1]
	return decrypt(input, key)
}

func (n *Notes) getKey(part int) (string, error) {
	cachePath := filepath.Join("cache", fmt.Sprintf("%d-%d-%d-keys.json", n.config.Event, n.quest, n.config.Seed))
	keys := map[string]string{}
	data, err := os.ReadFile(cachePath)
	if err == nil {
		if err := json.Unmarshal(data, &keys); err != nil {
			return "", fmt.Errorf("decode cached keys: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return "", fmt.Errorf("read cached keys: %w", err)
	}
	keyName := fmt.Sprintf("key%d", part)
	if keys[keyName] != "" {
		return keys[keyName], nil
	}
	if keys == nil {
		keys = map[string]string{}
	}
	if n.config.Session == "" {
		return "", fmt.Errorf("SESSION is required to retrieve decryption keys")
	}
	url := fmt.Sprintf("https://api.everybody.codes/event/%d/quest/%d", n.config.Event, n.quest)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	request.AddCookie(&http.Cookie{Name: "everybody-codes", Value: n.config.Session})
	client := &http.Client{Timeout: 30 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("retrieve keys: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("retrieve keys: HTTP %d", response.StatusCode)
	}
	var available struct {
		Key1 string `json:"key1"`
		Key2 string `json:"key2"`
		Key3 string `json:"key3"`
	}
	if err := json.NewDecoder(response.Body).Decode(&available); err != nil {
		return "", fmt.Errorf("decode keys: %w", err)
	}
	for name, key := range map[string]string{"key1": available.Key1, "key2": available.Key2, "key3": available.Key3} {
		if key != "" {
			keys[name] = key
		}
	}
	if len(keys) > 0 {
		data, err := json.Marshal(keys)
		if err != nil {
			return "", fmt.Errorf("encode cached keys: %w", err)
		}
		if err := os.MkdirAll(filepath.Dir(cachePath), 0700); err != nil {
			return "", fmt.Errorf("create cache directory: %w", err)
		}
		if err := os.WriteFile(cachePath, data, 0600); err != nil {
			return "", fmt.Errorf("cache keys: %w", err)
		}
	}
	if keys[keyName] == "" {
		return "", fmt.Errorf("part %d is locked: key unavailable", part)
	}
	return keys[keyName], nil
}

func decrypt(input, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("invalid decryption key: %w", err)
	}
	data, err := hex.DecodeString(input)
	if err != nil {
		return "", fmt.Errorf("decode encrypted input: %w", err)
	}
	if len(data) == 0 || len(data)%aes.BlockSize != 0 {
		return "", fmt.Errorf("invalid encrypted input length")
	}
	cipher.NewCBCDecrypter(block, []byte(key)[:aes.BlockSize]).CryptBlocks(data, data)
	padding := int(data[len(data)-1])
	if padding < 1 || padding > aes.BlockSize {
		return "", fmt.Errorf("invalid input padding")
	}
	for _, b := range data[len(data)-padding:] {
		if int(b) != padding {
			return "", fmt.Errorf("invalid input padding")
		}
	}
	data = data[:len(data)-padding]
	if !utf8.Valid(data) {
		return "", fmt.Errorf("decrypted input is not valid UTF-8")
	}
	return string(data), nil
}
