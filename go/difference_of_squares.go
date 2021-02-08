package diffsquares

// SquareOfSum returns the square of the sum till number n
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares returns the sum of squares of each number till n
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference provides the difference of SquareOfSum & SumOfSquares
// for a number n
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
