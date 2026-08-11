package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	tags := map[string]bool{
		"[TRC]": true,
		"[DBG]": true,
		"[INF]": true,
		"[WRN]": true,
		"[ERR]": true,
		"[FTL]": true,
	}

	re := regexp.MustCompile(`^\[[A-Z]{3}\]`)
	tag := re.FindString(text)

	if tag != "" {
		exists := tags[tag]

		return exists
	}

	return false
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-*~=]*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	count := 0

	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)

	tagged := make([]string, 0, len(lines))

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)

		if len(matches) == 2 {
			tagged = append(tagged, fmt.Sprintf("[USR] %s %s", matches[1], line))
		} else {
			tagged = append(tagged, line)
		}
	}

	return tagged
}
