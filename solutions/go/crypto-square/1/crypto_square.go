package cryptosquare

import (
	"regexp"
	"strings"
)

func Encode(pt string) string {
	normalized := Normalize(pt)
	cols, rows := ChunkSizes(normalized)
	composed := Compose(normalized, cols, rows)

	return composed
}

func Normalize(pt string) string {
	return strings.ToLower(regexp.MustCompile(`\W+`).ReplaceAllString(pt, ""))
}

func ChunkSizes(pt string) (int, int) {
	columns, rows, length := 0, 0, len(pt)

	for columns*rows < length {
		if columns == rows {
			columns++
		} else {
			rows++
		}
	}

	return columns, rows
}

func Compose(normalized string, cols, rows int) string {
	length := len(normalized)
	var composed strings.Builder

	for i := 0; i < cols; i++ {
		if i > 0 {
			composed.WriteByte(' ')
		}

		for j := 0; j < rows; j++ {
			index := j*cols + i
			if index <= length-1 {
				composed.WriteByte(normalized[index])
			} else {
				composed.WriteByte(' ')
			}
		}
	}

	return composed.String()
}
