	package luhn

	import (
		"regexp"
		"strconv"
		"strings"
	)

	// Valid ( s string ) returns bool values:
	//	- true if the given numerical string is a valid Identifier Number according to Luhn Algorithm
	//	- false if the given numerical string is not a valid Identifier Number.....
	func Valid( s string ) bool {

		var exp = regexp.MustCompile("^[0-9]+[0-9 ]*[0-9]+$")
		// The above RegExp verfies whether the given string starts with a number and ends with a number
		// allowing only white spaces and not special chars like :, #, %, $, ...
		if !exp.MatchString(s) {

			return false
		} 

		sum := 0
		s = strings.ReplaceAll(s, " ", "")
		ind := len(s) % 2
		// ind(indicator) indicates to double odd or even index string numerals....
		// if ind = 1 - then we have double the odd index(1, 3, 5, 7, 9....) string numerals...
		// if ind = 0 - then we have double the even index(0, 2, 4, 6, 8....) string numerals...
		for i := 0; i < len(s); i ++ {

			if i % 2 == ind {

				temp, _ := strconv.Atoi(string(s[i]))
				temp *= 2
				if temp >= 10 {

					temp -= 9
				}
				sum += temp
			} else {

				temp, _ := strconv.Atoi(string(s[i]))
				sum += temp
			}
		}
		if sum % 10 == 0 {
			
			return true	 
		} 
		return false
	}