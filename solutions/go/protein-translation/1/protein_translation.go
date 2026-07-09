package proteintranslation

import (
	"errors"
)

var codonToAminoAcid = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var ErrStop = errors.New("This is a STOP codon")
var ErrInvalidBase = errors.New("The codon is not valid")

func FromRNA(rna string) ([]string, error) {
	length := len(rna)
	found := []string{}

	if length < 3 {
		return found, nil
	}

	for i := 0; i < length; i = i + 3 {
		end := i + 3
		if end > length {
			end = length - 1
		}

		amino, err := FromCodon(rna[i:end])

		if err == ErrInvalidBase {
			return found, err
		}

		if err != nil {
			return found, nil
		}

		found = append(found, amino)
	}

	return found, nil
}

func FromCodon(codon string) (string, error) {
	length := len(codon)
	if length != 3 {
		return "", ErrInvalidBase
	}

	amino, exists := codonToAminoAcid[codon]
	if !exists {
		return "", ErrInvalidBase
	}

	if amino == "STOP" {
		return "", ErrStop
	}

	return amino, nil
}
