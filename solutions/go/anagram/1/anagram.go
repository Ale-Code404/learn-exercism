package anagram

import (
	"slices"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var anagrams []string

	subjectLower := strings.ToLower(subject)
	subjectSplit := strings.Split(subjectLower, "")
	slices.Sort(subjectSplit)

	subjectNormalized := strings.Join(subjectSplit, "")

	for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)
		if candidateLower == subjectLower {
			continue
		}

		normalized := strings.Split(strings.ToLower(candidate), "")
		slices.Sort(normalized)
		
		if strings.Join(normalized, "") == subjectNormalized {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}
