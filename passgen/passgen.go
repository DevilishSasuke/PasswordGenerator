package passgen

import (
	"math/rand"
	"time"
)

func GeneratePassword(length int, charset []byte) []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Int()

	password := make([]byte, length)

	for i := range password {
		password[i] = charset[r.Intn(len(charset))]
	}

	return password
}

func ComposeAlphabet(isDigit bool, isSpecChar bool) []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	specChars := "!@#$%^&*()-_=+[]{}|;:,.<>?"

	if isDigit {
		charset += digits
	}
	if isSpecChar {
		charset += specChars
	}

	sequence := []byte(charset)

	r.Shuffle(len(sequence), func(i, j int) {
		sequence[i], sequence[j] = sequence[j], sequence[i]
	})

	return sequence
}
