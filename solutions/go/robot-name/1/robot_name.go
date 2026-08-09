package robotname

import (
	"crypto/rand"
	"math/big"
)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	r.generateName()

	return r.name, nil
}

func (r *Robot) Reset() {
	r.generateName()
}

func (r *Robot) generateName() {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "1234567890"

	number := func(max int64) int {
		n, err := rand.Int(rand.Reader, big.NewInt(max))
		if err != nil {
			return 0
		}

		return int(n.Int64())
	}

	chars := []byte{
		letters[number(26)],
		letters[number(26)],
		digits[number(10)],
		digits[number(10)],
		digits[number(10)],
	}

	r.name = string(chars)
}
