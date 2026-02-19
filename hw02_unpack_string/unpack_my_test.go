package hw02unpackstring

import (
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
