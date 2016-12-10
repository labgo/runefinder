package main

import (
	"strconv"
	"strings"
)

func tokenize(s string) []string {
	separator := func(c rune) bool {
		return c == ' ' || c == '-'
	}
	return strings.FieldsFunc(s, separator)
}

func stringInSlice(needle string, haystack []string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

// Parse takes a line from the UCD (Unicode Character Database) text
// and returns: the rune, words from the name and legacy name
func Parse(ucdLine string) (rune, []string) {
	fields := strings.Split(ucdLine, ";")
	code64, _ := strconv.ParseInt(fields[0], 16, 0)
	uchar := rune(code64)
	words := tokenize(fields[1])
	if len(fields[10]) > 0 {
		for _, word := range tokenize(fields[10]) {
			if !stringInSlice(word, words) {
				words = append(words, word)
			}
		}
	}
	return uchar, words
}
