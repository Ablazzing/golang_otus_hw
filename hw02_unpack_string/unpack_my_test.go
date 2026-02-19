package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpackSecond(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "bye", expected: "bye"},
		{input: "hello", expected: "hello"},
		{input: "hel2o", expected: "hello"},
	}

	for _, tc := range tests {
		tc := tc //nolint:all
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidStringSecond(t *testing.T) {
	invalidStrings := []string{"1", "ka101", "ly2l22o"}
	for _, tc := range invalidStrings {
		tc := tc //nolint:all
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
