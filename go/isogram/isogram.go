package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(input string) bool {
	inputMap := make(map[rune]bool)
	inputString := strings.ToUpper(input)
	for _, inputChar := range inputString {
		if unicode.IsLetter(inputChar) {
			if inputMap[inputChar] {
				return false
			}
			inputMap[inputChar] = true
		}
	}
	return true
}
