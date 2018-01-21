// Package diffsquares calculates the diff
// between square of sums and sum of squares
// of a number
package diffsquares

// SquareOfSums for the square of sums
func SquareOfSums(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares for the sum of sqr
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += (i * i)
	}
	return sum
}

// Difference gets the diff between the 2 above results
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
