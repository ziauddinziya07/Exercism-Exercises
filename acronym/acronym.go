// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {

	res := ""

	/*
	SMART - Specific, Measurable, Achievable, Relevant, Timely
	TASER - Thomas A. Swift’s Electric Rifle
	SCUBA - Self-Contained Underwater Breathing Apparatus
	WASP - White Anglo-Saxon Protestant
	TRNT - The Road _Not_ Taken

	****** Usually Acronym conatains ',', '.', '_', '-', ' '(white-space),  etc.,
		   which make sense to form a Acronym.....*******
	*/
	s = strings.ReplaceAll(s, ",", " ")
	s = strings.ReplaceAll(s, ".", " ")
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ReplaceAll(s, "-", " ") 

	// After removing all special chars, we can split the Acronym sentence 

	temp := strings.Split(s, " ")
	for _, v := range temp {
		if v != "" {
			res += string(v[0])
		}
	}

	return strings.ToUpper(res)
}
