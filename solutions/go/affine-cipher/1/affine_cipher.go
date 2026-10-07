package affinecipher

import (
	"errors"
	"strings"
	"unicode"
)

// ProcessLetter is a function to decide for each letter when to ignore, process (encrypt or decrypt) and returns the index
// of the letter in the alphabet
type ProcessLetter = func(ignore, process bool, index int)

var NotCoPrimesErr = errors.New("The numbers aren't co-primes")

func Encode(text string, a, b int) (string, error) {
	var encoded strings.Builder
	chars := 0

	if !isCoPrimes(a, 26) {
		return "", NotCoPrimesErr
	}

	for _, letter := range normalizeText(text) {
		var new rune

		onProcessLetter(letter, func(ignore, process bool, index int) {
			if !ignore {
				new = letter

				if process {
					new = transformChar(index, a, b, true)
				}
			}
		})

		if new != 0 {
			encoded.WriteRune(new)
			chars++
		}

		if chars != 0 && chars%5 == 0 {
			encoded.WriteString(" ")
			chars = 0
		}
	}

	return strings.TrimRight(encoded.String(), " "), nil
}

func Decode(text string, a, b int) (string, error) {
	if !isCoPrimes(a, 26) {
		return "", NotCoPrimesErr
	}

	var decoded strings.Builder

	for _, letter := range normalizeText(text) {
		onProcessLetter(letter, func(ignore, process bool, index int) {
			if !ignore && process {
				decoded.WriteRune(transformChar(index, a, b, false))
			} else if !ignore {
				decoded.WriteRune(letter)
			}
		})
	}

	return decoded.String(), nil
}

func onProcessLetter(letter rune, on ProcessLetter) {
	isLetter := unicode.IsLetter(letter)
	isNumber := unicode.IsNumber(letter)

	isAllowed := isLetter || isNumber

	on(!isAllowed, isLetter, int(letter)-97)
}

func normalizeText(text string) string {
	return strings.ToLower(text)
}

func isCoPrimes(a, b int) bool {
	for b != 0 {
		a, b = b, a%b
	}

	return a == 1
}

func transformChar(index int, a, b int, encrypt bool) rune {
	if encrypt {
		return rune(97 + (a*index+b)%26)
	}

	// Decrypt: D(y) = a^-1 * (y - b) mod 26
	aInv := modInverse(a, 26)
	return rune(97 + (aInv*(index-b)%26+26)%26)
}

func modInverse(a, m int) int {
	for i := 1; i < m; i++ {
		if (a*i)%m == 1 {
			return i
		}
	}

	return 1
}
