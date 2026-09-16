package nthprime

import (
	"errors"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("input must be start from one")
	}

	on := 0
	prime := 1
	divs := 0

	for on != n {
		prime++

		for v := range prime {
			if prime%(v+1) == 0 {
				divs++
			}

			if divs > 2 {
				break
			}
		}

		if divs == 2 {
			on++
		}

		divs = 0
	}

	return prime, nil
}
