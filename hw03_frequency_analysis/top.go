package hw03frequencyanalysis

import (
	"maps"
	"sort"
	"strings"
)

func Top10(value string) []string {
	if strings.TrimSpace(value) == "" {
		return []string{}
	}

	frequencyMap := make(map[string]int)
	for _, word := range strings.Fields(value) {
		frequencyMap[word]++
	}

	words := make([]string, 0, len(frequencyMap))
	for value := range maps.Keys(frequencyMap) {
		words = append(words, value)
	}

	comparingWords := func(i, j int) bool {
		word1 := words[i]
		word2 := words[j]

		return frequencyMap[word1] > frequencyMap[word2] ||
			(frequencyMap[word1] == frequencyMap[word2] && word1 < word2)
	}

	sort.Slice(words, comparingWords)
	return words[0:10]
}
