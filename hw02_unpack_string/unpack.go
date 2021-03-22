package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(raw string) (string, error) {
	if raw == "" {
		return "", nil
	}

	if unicode.IsDigit(rune(raw[0])) {
		return "", ErrInvalidString
	}

	var prevChar rune
	chars := make([]string, len(raw))

	for _, char := range raw {
		if unicode.IsDigit(char) {
			if unicode.IsDigit(prevChar) {
				return "", ErrInvalidString
			}

			repeats, _ := strconv.Atoi(string(char))
			switch {
			case repeats > 1:
				chars = append(chars, strings.Repeat(string(prevChar), repeats-1))
			case repeats == 0:
				chars = chars[:len(chars)-1]
			case repeats < 0:
				return "", ErrInvalidString
			}

			prevChar = char
			continue
		}

		chars = append(chars, string(char))
		prevChar = char
	}

	return strings.Join(chars, ""), nil
}
