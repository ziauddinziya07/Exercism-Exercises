// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"regexp"
	)

func check( s string ) bool {
	return strings.ContainsAny(s, "1234567890/*-+~_()=@#$%^&[]{}")
}

// Hey should have a comment documenting it.
func Hey(remark string) ( res string ) {

	// ************ Iteration - 3 ***********
	remark = strings.TrimSpace(remark)
	res = "Fine. Be that way!"

	if remark != "" {
		
		exp := regexp.MustCompile("^([0-9]+[,./*-+~_()=@#$%^&{} ]*)+$")  	// To check whether a given string is of type "1-12-1" or "1, 2, 3...." 
		isYelling := remark == strings.ToUpper(remark)
		isQues := remark[len(remark) - 1] == '?'
		isAlpNum := check(remark)
		isNumeric := exp.MatchString(remark)  // To check whether the string contains only numeric-values.....

		if isYelling  &&  !isAlpNum  &&  isQues { 
			// The string must not be a Numeric and AlphaNumeric.....
			res = "Calm down, I know what I'm doing!"

		} else if isYelling && !isQues && !isNumeric {
			// The String must not Numeric but it can be AlphaNumeric....
			res = "Whoa, chill out!"
		
		} else if isQues {
			// Whatever String may be, It must be a Question.....
			res = "Sure."

		} else {

			res = "Whatever."

		}
	}	
	
	return
}


/*

	************ Iteration - 1 **************
	remark = strings.TrimSpace(remark)
	
	if remark == "" {
		return "Fine. Be that way!"
	} else if remark == strings.ToUpper(remark) && string(remark[len(remark) - 1]) == "?" && !strings.ContainsAny(remark, "4:)") {
		return "Calm down, I know what I'm doing!"
	} else if string(remark[len(remark) - 1]) == "?" {
		return "Sure."
	} else if remark == strings.ToUpper(remark) && strings.ContainsAny(remark, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return "Whoa, chill out!"	
	} else {
		return "Whatever."
	}


	************ Iteration - 2 **************
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	} else if remark == strings.ToUpper(remark) && string(remark[len(remark)-1]) == "?" && !strings.ContainsAny(remark, "0123456789:)") {
		return "Calm down, I know what I'm doing!"
	} else if string(remark[len(remark)-1]) == "?" {
		return "Sure."
	} else if remark == strings.ToUpper(remark) && strings.ContainsAny(remark, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}

*/