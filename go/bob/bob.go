package bob

import (
	"strings"
	"unicode"
)

// Remark is a convenience type for identifying what kind of remark it is.
type Remark struct {
	remark string
}

func (r Remark) isSilence() bool {
	return r.remark == ""
}

func (r Remark) isShouting() bool {
	// IndexFuncが引数に取る関数はIsLetterのシグネチャと同じだからIsLetterが引数になっているのだろうが、
	// 結局、 IsLetterを満たす最初の文字のインデックスを返すのか
	// つまり、文字が一つも含まれていない場合は-1を返すから↓のようにして文字が含まれているかを調べているのだ。
	// IndexFunc内で引数の関数、つまりIsLetterに引数を私ているから↓で省略できるのだろう。
	hasLetters := strings.IndexFunc(r.remark, unicode.IsLetter) >= 0
	isUpcased := strings.ToUpper(r.remark) == r.remark

	return hasLetters && isUpcased
}

func (r Remark) isQuestion() bool {
	return strings.HasSuffix(r.remark, "?")
}

func (r Remark) isExasperated() bool {
	return r.isShouting() && r.isQuestion()
}

func newRemark(remark string) Remark {
	return Remark{strings.TrimSpace(remark)}
}

// Hey greets with an appropriate remark.
func Hey(remark string) string {
	r := newRemark(remark)
	switch {
	case r.isSilence():
		return "Fine. Be that way!"
	case r.isExasperated():
		return "Calm down, I know what I'm doing!"
	case r.isShouting():
		return "Whoa, chill out!"
	case r.isQuestion():
		return "Sure."
	default:
		return "Whatever."
	}
}