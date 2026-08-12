package piglatin

import (
	"fmt"
	"regexp"
	"strings"
)

func Sentence(sentence string) string {
	var translation strings.Builder

	words := strings.Split(sentence, " ")
	length := len(words)

	vowels := "aeiou"
	consonants := "bcdfghjklmnpqrstxyz"

	rule1 := regexp.MustCompile(fmt.Sprintf(`^([%s]|xr|yt)`, vowels))
	rule2 := regexp.MustCompile(fmt.Sprintf(`^([%s]+)(\w+)`, consonants))
	rule3 := regexp.MustCompile(fmt.Sprintf(`^([%s]*qu)(\w+)`, consonants))
	rule4 := regexp.MustCompile(fmt.Sprintf(`^([%s]+)(y\w+)`, consonants))

	for i := 0; i < length; i++ {
		word := words[i]
		translated := word
		matched := false

		if rule1.MatchString(word) && !matched {
			translated = word + "ay"
			matched = true
		}

		if rule3.MatchString(word) && !matched {
			match := rule3.FindAllStringSubmatch(word, -1)
			if match != nil {
				translated = match[0][2] + match[0][1] + "ay"
				matched = true
			}
		}

		if rule4.MatchString(word) && !matched {
			match := rule4.FindAllStringSubmatch(word, -1)
			if match != nil {
				translated = match[0][2] + match[0][1] + "ay"
				matched = true
			}
		}

		if rule2.MatchString(word) && !matched {
			match := rule2.FindAllStringSubmatch(word, -1)
			if match != nil {
				translated = match[0][2] + match[0][1] + "ay"
				matched = true
			}
		}

		if length > 1 && i > 0 {
			translation.WriteString(fmt.Sprintf(" %s", translated))
		} else {
			translation.WriteString(translated)
		}
	}

	return translation.String()
}
