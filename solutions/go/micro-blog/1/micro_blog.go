package microblog

func Truncate(phrase string) string {
	length := 0
	limit := 0

	for index := range phrase {
		length++
		limit = index

		if length == 6 {
			break
		}
	}

	if length < 6 {
		return phrase
	}

	return phrase[:limit]
}
