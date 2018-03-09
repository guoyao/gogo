package x_strings

import (
	"strings"
)

func IndexOf(str, substr string) int {
	asciiPos := strings.Index(str, substr)
	if asciiPos <= 0 {
		return asciiPos
	}

	pos, totalSize, reader := 0, 0, strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize, pos = totalSize+size, pos+1
		if totalSize == asciiPos {
			return pos
		}
	}

	return pos
}
