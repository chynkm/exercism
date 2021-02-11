package grains

import (
	"errors"
)

const numberOfSquares = 64

// Square returns the number of wheat in a chessboard square
func Square(n int) (uint64, error) {
	if n > numberOfSquares || 1 > n {
		return 0, errors.New("square value should be between 1 and 64")
	}

	// n << x is "n times 2, x times". And y >> z is "y divided by 2, z times".
	// For example, 1 << 5 is "1 times 2, 5 times" or 32.
	// And 32 >> 5 is "32 divided by 2, 5 times" or 1.
	return 1 << (n - 1), nil
}

// Total returns the total grains in the chessboard
func Total() uint64 {
	total, err := Square(64)

	if err != nil {
		panic("got error")
	}

	return total*2 - 1
}
