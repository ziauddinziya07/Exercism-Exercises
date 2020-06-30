package reverse


func Reverse ( s string ) ( r string ) {

	// ************ It is not a common thing to reverse a string in GO like other languages!!!!!
	// 		for i := len(s) - 1; i >= 0; i -- {
	// 				r += string(s[i])
	//		 } 
	// 				Due to Golang! supports Unicode UTF - 8, The test cases might contains strings from other languages....
	//				Eg. "Hello, 世界", "غزالہ and ضياء"
	//				So to Handle the letters of different languages Go uses Unicode
	//				In Go, Unicodes are stored in runes (i.e. int32)				
	
	var temp = []rune(s)
	for i := len(temp) - 1; i >= 0; i -- {

		r += string(temp[i])
	}	
	return
}

