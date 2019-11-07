package reverse

import (
	"bytes"
)

func Reverse(s string) string {
	var ret bytes.Buffer

	runes := []rune(s)
	for i := len(runes)-1; i >= 0;  i-- {
		ret.WriteRune(runes[i])
	}
	return ret.String()
}