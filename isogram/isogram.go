package isogram

import (
	"regexp"
	"strings"
)

func IsIsogram(s string) bool {

	// This Exercise is based on Sequences and Strings, so we use a RegExp to match the sequence.....
	if len(s) > 1 {

		s = strings.ToUpper(s) // The given string might contain combination of Lower and Upper case Alphabets....
		for _, v := range s {

			if    ( strings.Count(s, string(v)) > 1 )   &&   regexp.MustCompile("[A-Z]").MatchString(string(v)) {
				// The each letter of string must contain only one occurance in the string, so that the String is called ** Isogram **
				// Each letter of string must be from the set of [A-Z]
				return false
			}
		}
	}
	return true
}
