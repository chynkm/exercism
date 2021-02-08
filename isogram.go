package isogram

import "strings"

// IsIsogram checks whether a string is an isogram
func IsIsogram(word string) bool {
	seen := map[rune]bool{}

	for _, r := range strings.ToUpper(word) {
		if r == '-' || r == ' ' {
			continue
		}

		if seen[r] {
			return false
		}

		seen[r] = true
	}

	return true
}
