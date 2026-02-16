package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(value string) (string, error) {
	builder := strings.Builder{}
	previousLetter := ""
	for _, value := range value {
		if letterCount, err := strconv.Atoi(string(value)); err == nil {
			if previousLetter == "" {
				return "", ErrInvalidString
			}
			builder.WriteString(strings.Repeat(previousLetter, letterCount))
			previousLetter = ""
			continue
		}
		if previousLetter != "" {
			builder.WriteString(previousLetter)
		}
		previousLetter = string(value)
	}
	if previousLetter != "" {
		builder.WriteString(previousLetter)
	}
	return builder.String(), nil
}
