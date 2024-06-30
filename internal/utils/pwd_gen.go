package utils

import (
	"crypto/rand"
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

func GeneratePassword(length int, useUpper, useLower, useSpecial, useNumbers bool) string {
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

	var password strings.Builder

	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		password.WriteByte(chars[randomIndex.Int64()])
	}
	return password.String()
}

func CopyToClipboard(text string) error {
	err := clipboard.WriteAll(text)
	return err
}