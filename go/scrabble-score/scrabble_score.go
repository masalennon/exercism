package scrabble

import (
	"strings"
)

var scoreMap = make(map[rune]int)

func init() {
	classes := map[string]int{
		"AEIOULNRST": 1,
		"DG":         2,
		"BCMP":       3,
		"FHVWY":      4,
		"K":          5,
		"JX":         8,
		"QZ":         10,
	}
	for char, score := range classes {
		for _, letter := range char {
			scoreMap[letter] = score
		}
	}
}

func Score(word string) (score int) {

	word = strings.ToUpper(word)
	for _, c := range word {
		score += scoreMap[c]
	}
	return
}
