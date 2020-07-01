package diffsquares

import "math"

// SquareOfSum ( n int ) int will first calculate the sum of first n Natural Numbers
// i.e. 1 + 2 + 3 + ..... + n and then return square of sum....
func SquareOfSum( n int ) int {

	sum := 0.0
	for i := 1.0; i <= float64(n); i++ {

		sum += i
	}
	return int(math.Pow(sum, 2))
}

// SumOfSquares ( n int ) int returns the sum of squares of first n Natural Numbers.....
func SumOfSquares( n int ) int {

	sum := 0.
	for i := 1.; i <= float64(n); i ++ {

		sum += math.Pow(i, 2)
	}
	return int(sum)
}

// Difference ( n int ) int returns the difference between Squares of Sum and Sum of Squares of first n Natural Numbers.....
func Difference( n int ) int {

	return SquareOfSum(n) - SumOfSquares(n)
} 