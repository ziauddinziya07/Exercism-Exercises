package scrabble

import "strings"

func Score ( str string ) ( x int ) {
	// This Exercise is on For loop, Maps and strings.......
	// So thats why we use Maps...
	str = strings.ToUpper(str) 	// String might be in lower so we convert it to Upper - case
	scores := map[string]int {

		"AEIOULNRST": 1,
		"DG": 2,
		"BCMP": 3,
		"FHVWY": 4,
		"K": 5,
		"JX": 8,
		"QZ": 10,
	}

	for _, s := range str {

		for k, v := range scores {

			if strings.ContainsAny(k, string(s)) {
				x += v
			}
		}
	}
	return
}