package grains

import (
	"errors"
	"math"
)

const numberOfSquares = 64

// Square returns the number of wheat in a chessboard square
func Square(n int) (uint64, error) {
	if n > numberOfSquares || 1 > n {
		return 0, errors.New("square value should be between 1 and 64")
	}

	return uint64(math.Pow(2, float64(n)-1)), nil
}

// Total returns the total grains in the chessboard
func Total() uint64 {
	var total uint64

	for i := 1; i <= numberOfSquares; i++ {
		n, _ := Square(i)
		total += n
	}

	return total
}
