

import "fmt"


func Proverb(rhyme []string) (res []string) {

	if len(rhyme) != 0 {

		if len(rhyme) > 1 {
			
			for i := 0; i < len(rhyme)-1; i++ {
				res = append(res, fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i], rhyme[i+1]))
			}
		}
		res = append(res, fmt.Sprintf("And all for the want of a %v.", rhyme[0]))
	}
	return
}
