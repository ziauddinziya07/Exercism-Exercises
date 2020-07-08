package grains

import (
	"fmt"
)

// Square function returns the no. of grains present on the nth block of the Chess Board ....
func Square(n int) (uint64, error) {

	if n < 1 || n > 64 {

		return 0, fmt.Errorf("The %dth block does not exist in the Chess Board", n)
	}

	// Calculating the result using Power function will give result at the time-complexity of O(n) ....
	// return pow(2, n - 1), nil
	// return uint64(math.Pow(2, float64(n - 1))), nill

	// This will give the result 2**(n-1) in just one operation ....
	// I observed that code optimized from O(n) ---> O(1) ....
	return 1 << (n - 1), nil
}

// Total returns the total no. of grains present on the chess from 1st block to 64th block ....
func Total() uint64 {

	return 1 << 64 - 1
}

// pow function return the x**n
// func pow(x, n int) uint64 {

// 	var r uint64 = 1
// 	for i := 1; i <= n; i ++ {

// 		r *= uint64(x)
// 	}
// 	return r
// }