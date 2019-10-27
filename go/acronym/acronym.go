// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"fmt"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	//まずは記号を取り除く
	//スペース区切りで単語をスライスに格納し、それぞれの単語の一番最初の文字を連結させた文字列を作る
	charToReplace := []string{"-", ",", "_"}
	replacedStr := replaceChar(s, charToReplace)
	fmt.Println(replacedStr)

	words := strings.Fields(replacedStr)
	abb := ""
	for i := range words {
		abb += string(words[i][0])
		fmt.Println(string(words[i][0]))
	}
	return strings.ToUpper(abb)
}

//文字列とそこから除去したい記号をスライスで渡すと、それを除去した文字列を返してくれる
func replaceChar(s string, c []string) (str string) {
	str = s
	for _, i := range c {
		str = strings.Replace(str, i, " ", -1)
	}
	return
}
