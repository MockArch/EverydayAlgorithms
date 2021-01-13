// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"fmt"
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	s = strings.TrimSpace(s)
	var outPut string
	for i, w := range s {

		if i == 0 {
			fmt.Println(string(w))
			outPut += string(w)
		} else if !unicode.IsLetter(rune(s[i-1])) && rune(s[i-1]) != 39 {
			if unicode.IsLetter(w) {
				outPut += strings.ToUpper(string(w))
			}

		}

	}

	return outPut
}
