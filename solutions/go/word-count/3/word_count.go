package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	re := regexp.MustCompile(`[A-Za-z0-9]+('[A-Za-z0-9]+)?`)
	freq := map[string]int{}

	matches := re.FindAllString(strings.ToLower(phrase), -1)

	for _, word := range matches {
		_, exist := freq[word]
		if !exist {
			freq[word] = 1
		} else {
			freq[word]++
		}
	}

	return freq
}
