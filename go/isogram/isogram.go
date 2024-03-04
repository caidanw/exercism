package isogram

import (
	"slices"
	"strings"
)

var ignoredRunes = []rune(" -")

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	letters := make(map[rune]bool, len(word))

	for _, token := range word {
		if slices.Contains(ignoredRunes, token) {
			continue
		}

		if _, ok := letters[token]; ok {
			return false
		}

		letters[token] = true
	}

	return true
}
