// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

func check( s string ) bool {
	return strings.ContainsAny(s, "1234567890/*-+~_()=@#$%^&[]{}")
}

// Hey should have a comment documenting it.
func Hey(remark string) ( res string ) {

	remark = strings.TrimSpace(remark)

	if remark == "" {
		res = "Fine. Be that way!"
	} else {

		isYelling := remark == strings.ToUpper(remark)
		isQues := remark[len(remark) - 1] == '?'
		isAlpNum := check(remark)

		if isYelling  &&  !isAlpNum  &&  isQues {

			res = "Calm down, I know what I'm doing!"

		} else if 
	}
	
	
	
	else if remark == strings.ToUpper(remark) && string(remark[len(remark)-1]) == "?" && !strings.ContainsAny(remark, "0123456789:)") {
		return "Calm down, I know what I'm doing!"
	} else if string(remark[len(remark)-1]) == "?" {
		return "Sure."
	} else if remark == strings.ToUpper(remark) && strings.ContainsAny(remark, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}
	return
}
