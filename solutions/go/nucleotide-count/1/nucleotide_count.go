package nucleotidecount

import (
	"errors"
	"regexp"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[byte]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram = map[byte]int{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	exp := regexp.MustCompile("^[ACGT]+$")
	if !exp.MatchString(string(d)) && len(d) != 0 {
		return h, errors.New("Invalid strand")
	}

	for _, ncl := range d {
		h[byte(ncl)]++
	}

	return h, nil
}
