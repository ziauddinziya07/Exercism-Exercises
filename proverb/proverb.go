// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

var (
	lastSen string = "And all for the want of a "
	fstSen string = "For want of a "
	secSen string = " the "
	thrdSen string = " was lost."
)
// Proverb should have a comment documenting it.
func Proverb(rhyme []string) (res []string) {
	
	if len(rhyme) != 0  {
		
		if len(rhyme) > 1 {
			for i := 0; i < len(rhyme) - 1; i ++ {
				res = append(res, fstSen + rhyme[i] + secSen + rhyme[i + 1] + thrdSen)
			}
		}
		res = append(res, lastSen + rhyme[0] + ".")		
	}
	return
}
