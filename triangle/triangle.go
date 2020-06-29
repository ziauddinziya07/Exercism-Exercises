

package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
// ********* Here the triangle_test.go file is expecting result of type "Kind"
//			 So we used a custom-type ************
type Kind string

const (
	
    NaT, Equ, Iso, Sca Kind = "NaT", "Equ", "Iso", "Sca"
)

// ***** The function checkVar() evaluates whether
// 		- The given number is greater than 0
//		- The given number is b/w 0 and +Inf 
func checkVar( x float64 ) bool {

	if x > 0 && x != math.Inf(-1) && x != math.Inf(1) {

		return true
	}
	return false
}

func KindFromSides(a, b, c float64) Kind {
	
	var k Kind = NaT
	// Evaluating: The given number is valid or not for our cases..... 
	if checkVar(a) && checkVar(b) && checkVar(c) {
		// Evaluating: The given Sides form a triangle........
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
