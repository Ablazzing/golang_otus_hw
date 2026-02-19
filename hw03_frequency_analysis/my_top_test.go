package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var textSecond = `банан банан банан арбуз арбуз арбуз велосипед велосипед яд яд глина дверь енот жизнь кот мотоцикл`

func TestTop10Second(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		expected := []string{
			"арбуз",     // 3
			"банан",     // 3
			"велосипед", // 2
			"яд",        // 2
			"глина",     // 1
			"дверь",     // 1
			"енот",      // 1
			"жизнь",     // 1
			"кот",       // 1
			"мотоцикл",  // 1
		}
		require.Equal(t, expected, Top10(textSecond))
	})
}
