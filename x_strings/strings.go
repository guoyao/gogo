package x_strings

import (
	"strings"
	"unicode/utf8"
)

func IsBlank(str string) bool {
	return strings.Trim(str, " ") == ""
}

// Rune Len
func Len(str string) int {
	return utf8.RuneCountInString(str)
}

// Rune At
func At(str string, index int) string {
	i, reader := 0, strings.NewReader(str)
	if index < 0 {
		index = Len(str) + index
	}

	for s, _, err := reader.ReadRune(); err == nil; s, _, err = reader.ReadRune() {
		if i == index {
			return string(s)
		}
		i++
	}

	return ""
}

func Index(str, substr string) int {
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
