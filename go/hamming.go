// Package hamming implements the Hamming distance calculation.
package hamming

import (
	"errors"
)

// Distance calculates the Hamming distance of two equal length strings.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("the strings are of different lengths")
	}

	runeA := []rune(a)
	runeB := []rune(b)
	distance := 0

	for i := 0; i < len(runeA); i++ {
		if runeA[i] != runeB[i] {
			distance++
		}
	}

	return distance, nil
}
