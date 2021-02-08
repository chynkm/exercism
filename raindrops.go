package raindrops

import (
	"strconv"
)

const (
	threeDrop = "Pling"
	fiveDrop  = "Plang"
	sevenDrop = "Plong"
)

func Convert(number int) string {
	var raindrop string

	if number%3 == 0 {
		raindrop += threeDrop
	}

	if number%5 == 0 {
		raindrop += fiveDrop
	}

	if number%7 == 0 {
		raindrop += sevenDrop
	}

	if len(raindrop) > 0 {
		return raindrop
	}

	return strconv.Itoa(number)
}
