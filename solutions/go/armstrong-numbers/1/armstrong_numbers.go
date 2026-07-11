package armstrongnumbers

import (
	"math"
	"strconv"
	"strings"
)

func IsNumber(n int) bool {
	str := strconv.Itoa(n)
	sum := 0

	digits := strings.Split(str, "")
	length := len(digits)

	for _, digit := range digits {
		value, err := strconv.ParseInt(digit, 10, 8)
		if err == nil {
			sum += int(math.Pow(float64(value), float64(length)))
		}
	}

	return sum == n
}
