package main

import (
	"strconv"
	"strings"
	"bufio"
	"os"
	"fmt"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func tokenize(s string) []string {
	separator := func(c rune) bool {
		return c == ' ' || c == '-'
	}
	return strings.FieldsFunc(s, separator)
}

func sliceHasString(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func sliceHasAllStrings(haystack []string, needles []string) bool {
	for _, needle := range needles {
		if !sliceHasString(haystack, needle) {
			return false
		}
	}
	return true
}

// Parse takes a line from the UCD (Unicode Character Database) text
// and returns: the rune, words from the name and legacy name
func Parse(ucdLine string) (rune, string, []string) {
	fields := strings.Split(ucdLine, ";")
	code64, _ := strconv.ParseInt(fields[0], 16, 0)
	uchar := rune(code64)
	name := fields[1]
	words := tokenize(name)
	if len(fields[10]) > 0 {
		appended := false
		for _, word := range tokenize(fields[10]) {
			if !sliceHasString(words, word) {
				words = append(words, word)
				appended = true
			}
		}
		if appended {
			name += " | " + fields[10]
		}
	}
	return uchar, name, words
}

func main() {
	var query []string
	if len(os.Args[1:]) > 0 {
		for _, word := range os.Args[1:] {
			query = append(query, strings.ToUpper(word))
		}
	} else {
		fmt.Println("Usage:  runefinder <word>...\texample: runefinder cat face")
		os.Exit(1)
	}
	ucd, err := os.Open("UnicodeData.txt")
  check(err)
	defer ucd.Close()
	input := bufio.NewScanner(ucd)
	for input.Scan() {
		uchar, name, words := Parse(input.Text())
		if sliceHasAllStrings(words, query) {
			fmt.Printf("%c\t%s\n", uchar, name)
		}

	}
}
