package notes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestDecode(t *testing.T) {
	t.Chdir(t.TempDir())
	key := "0123456789abcdef0123456789abcdef"
	input := "1,2,3\n"
	data := []byte(input)
	padding := aes.BlockSize - len(data)%aes.BlockSize
	for i := 0; i < padding; i++ {
		data = append(data, byte(padding))
	}
	block, _ := aes.NewCipher([]byte(key))
	cipher.NewCBCEncrypter(block, []byte(key)[:aes.BlockSize]).CryptBlocks(data, data)
	original := http.DefaultTransport
	defer func() { http.DefaultTransport = original }()
	requests := 0
	unlocked := false
	http.DefaultTransport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		requests++
		if r.URL.String() != "https://api.everybody.codes/event/2026/quest/1" {
			t.Errorf("unexpected URL: %s", r.URL)
		}
		cookie, err := r.Cookie("everybody-codes")
		if err != nil || cookie.Value != "test-token" {
			t.Error("missing session cookie")
		}
		body := fmt.Sprintf(`{"key1":%q,"answer1":"123"}`, key)
		if unlocked {
			body = fmt.Sprintf(`{"key2":%q}`, key)
		}
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body))}, nil
	})
	n := Notes{Part1: hex.EncodeToString(data), config: &Config{Event: 2026, Session: "test-token"}, quest: 1}
	got, err := n.Decode(1)
	if err != nil || got != input {
		t.Fatalf("Decode(1) = %q, %v", got, err)
	}
	if n.Part1 != hex.EncodeToString(data) {
		t.Error("Decode modified encoded notes")
	}
	// A new Notes instance should reuse the key stored on disk.
	cached := Notes{Part1: n.Part1, config: n.config, quest: 1}
	if got, err := cached.Decode(1); err != nil || got != input || requests != 1 {
		t.Fatalf("cached Decode(1) = %q, %v; requests = %d", got, err, requests)
	}
	if _, err = n.Decode(2); err == nil {
		t.Error("expected locked part error")
	}
	unlocked = true
	n.Part2 = n.Part1
	if got, err := n.Decode(2); err != nil || got != input || requests != 3 {
		t.Fatalf("unlocked Decode(2) = %q, %v; requests = %d", got, err, requests)
	}
	if _, err := cached.Decode(1); err != nil || requests != 3 {
		t.Fatalf("previously cached key lost: %v; requests = %d", err, requests)
	}
	if _, err = n.Decode(0); err == nil {
		t.Error("expected invalid part error")
	}
}

func TestDecryptInvalidInput(t *testing.T) {
	for _, input := range []string{"", "zz", "00", strings.Repeat("00", 16)} {
		if _, err := decrypt(input, "0123456789abcdef0123456789abcdef"); err == nil {
			t.Errorf("accepted invalid input %q", input)
		}
	}
}
