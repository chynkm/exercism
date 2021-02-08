package luhn

import (
	"strconv"
	"strings"
)

// Valid verifies whether a number is a LUHN or not
func Valid(num string) bool {
	num = strings.Replace(num, " ", "", -1)
	if len(num) <= 1 {
		return false
	}

	var sum int
	for flag, i := 0, len(num)-1; i >= 0; i, flag = i-1, flag+1 {
		n, err := strconv.Atoi(string(num[i]))
		if err != nil {
			return false
		}

		if flag%2 != 0 {
			d := n + n
			if d > 9 {
				d -= 9
			}
			sum += d
			continue
		}

		sum += n
	}

	return sum%10 == 0
}
