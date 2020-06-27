// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
    // Pick values for the following identifiers used by the test program.
    NaT, Equ, Iso, Sca Kind = "NaT", "Equ", "Iso", "Sca"
)

func checVar( x float64 ) bool {
	if x != 0 && x > math.Inf(-1) && x < math.Inf(1) {
		return true
	}
	return false
}
// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	
	var k Kind = NaT
	if checVar(a) && checVar(b) && checVar(c) {
		if (a + b >= c) && (a + c >= b ) && (b + c >= a) {
			switch {
				case a == b && b == c: k = Equ
				case (a == b && b != c) || (a == c && b != c) || (c == b && b != a): k = Iso
				default: k = Sca		
			}
		}
	}
	return k
}
