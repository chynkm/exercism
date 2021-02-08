package scrabble

import "strings"

// Score returns the scrabble value for the word
func Score(word string) int {
	score := 0
	for _, letter := range strings.ToUpper(word) {
		switch letter {
		case 'K':
			score += 5
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score++
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'J', 'X':
			score += 8
		case 'Q', 'Z':
			score += 10
		}
	}

	return score
}
