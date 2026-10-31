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
	http.DefaultTransport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if r.URL.String() != "https://api.everybody.codes/event/2026/quest/1" {
			t.Errorf("unexpected URL: %s", r.URL)
		}
		cookie, err := r.Cookie("everybody-codes")
		if err != nil || cookie.Value != "test-token" {
			t.Error("missing session cookie")
		}
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(fmt.Sprintf(`{"key1":%q,"answer1":"123"}`, key)))}, nil
	})
	n := Notes{Part1: hex.EncodeToString(data), config: &Config{Event: 2026, Session: "test-token"}, quest: 1}
	got, err := n.Decode(1)
	if err != nil || got != input {
		t.Fatalf("Decode(1) = %q, %v", got, err)
	}
	if n.Part1 != hex.EncodeToString(data) {
		t.Error("Decode modified encoded notes")
	}
	if _, err = n.Decode(2); err == nil {
		t.Error("expected locked part error")
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
