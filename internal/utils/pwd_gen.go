package utils

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"

	"github.com/atotto/clipboard"
)

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars  = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

func GeneratePassword(length int, useUpper, useLower, useSpecial, useNumbers bool) (string, error) {
	var chars string
	if useLower {
		chars += lowerChars
	}
	if useUpper {
		chars += upperChars
	}
	if useSpecial {
		chars += specialChars
	}
	if useNumbers {
		chars += numberChars
	}

	// Check if chars is empty
	if len(chars) == 0 {
		return "", errors.New("no character sets selected")
	}

	var password strings.Builder

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		password.WriteByte(chars[randomIndex.Int64()])
	}
	return password.String(), nil
}

// CopyToClipboard copies the provided text to the clipboard.
func CopyToClipboard(text string) error {
	return clipboard.WriteAll(text)
}
