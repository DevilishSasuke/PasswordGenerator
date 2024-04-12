package passgen

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

const (
	lower    = "abcdefghijklmnopqrstuvwxyz"
	upper    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits   = "0123456789"
	specials = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

func GeneratePassword(length int, charset *[]byte) ([]byte, error) {
	if length < 12 {
		return nil, fmt.Errorf("length is too short, should be >= 12, was %d", length)
	}

	password := make([]byte, length)

	for !IsStrong(&password, charset) {
		CreateCombination(&password, charset)
	}

	return password, nil
}

// Create a candidate for password
func CreateCombination(password *[]byte, charset *[]byte) {
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63n(1000)))

	for i := range *password {
		(*password)[i] = (*charset)[r.Intn(len(*charset))]
	}
}

// Create list of chosen characters
func ComposeAlphabet(hasDigits bool, hasSpecials bool) []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	charset := lower + upper

	if hasDigits {
		charset += digits
	}
	if hasSpecials {
		charset += specials
	}

	sequence := []byte(charset)

	r.Shuffle(len(sequence), func(i, j int) {
		sequence[i], sequence[j] = sequence[j], sequence[i]
	})

	return sequence
}

// Check if password is strong enough
func IsStrong(password *[]byte, charset *[]byte) bool {
	border := int(float64(len(*password)) * 0.3)

	if NumOfChars(password, lower) == 0 ||
		NumOfChars(password, upper) == 0 {
		return false
	}
	if slices.Contains(*charset, digits[0]) {
		digNum := NumOfChars(password, digits)
		if digNum == 0 || digNum > border {
			return false
		}
	}
	if slices.Contains(*charset, specials[0]) {
		specLen := NumOfChars(password, specials)
		if specLen == 0 || specLen > border {
			return false
		}
	}

	if NumberOfReps(password, 2) > 1 ||
		NumberOfReps(password, 3) > 0 {
		return false
	}

	return true
}

// return amount of ListOfChars elements in password
func NumOfChars(password *[]byte, listOfChars string) int {
	counter := 0

	for i := range listOfChars {
		if slices.Contains(*password, listOfChars[i]) {
			counter++
		}
	}

	return counter
}

// Check if sequence has size identical symbols in a row
// Size - length of subsequence to check
func NumberOfReps(password *[]byte, size int) int {
	if size < 2 {
		return len(*password)
	}
	counter := 0

	for i := 0; i < len(*password)-size; i++ {
		flag := true
		for j := i + 1; j < i+size; j++ {
			if (*password)[i] != (*password)[j] {
				flag = false
				break
			}
		}
		if flag {
			counter++
		}
	}

	return counter
}
