package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const ucdFileName = "UnicodeData.txt"
const ucdBaseURL = "http://www.unicode.org/Public/UCD/latest/ucd/"

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

// PrepareQuery takes a slice of strings and returns a slice of
// uppercased tokens, split at hyphens
func PrepareQuery(parts []string) []string {
	var query []string
	for _, part := range parts {
		for _, token := range tokenize(part) {
			query = append(query, strings.ToUpper(token))
		}
	}
	return query
}

func downloadUcdFile() {
	url := ucdBaseURL + ucdFileName
	fmt.Printf("%s not found\ndownloading %s\n", ucdFileName, url)
	running := make(chan bool)
	progressDisplay := func(running <-chan bool) {
		for {
			select {
			case <-running:
				fmt.Println("done!")
			case <-time.After(200 * time.Millisecond):
				fmt.Print(".")
			}
		}
	}
	go progressDisplay(running)
	defer func() {
		running <- false
	}()
	response, err := http.Get(url)
	check(err)
	defer response.Body.Close()
	file, err := os.Create(ucdFileName)
	check(err)
	_, err = io.Copy(file, response.Body)
	check(err)
	file.Close()
}

func getUcdFile() (*os.File, error) {
	ucd, err := os.Open(ucdFileName)
	if os.IsNotExist(err) {
		downloadUcdFile()
		ucd, err = os.Open(ucdFileName)
	}
	return ucd, err
}

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Usage:  runefinder <word>...\texample: runefinder cat face")
		os.Exit(1)
	}
	query := PrepareQuery(os.Args[1:])
	ucd, err := getUcdFile()
	check(err)
	defer ucd.Close()
	input := bufio.NewScanner(ucd)
	for input.Scan() {
		uchar, name, words := Parse(input.Text())
		if sliceHasAllStrings(words, query) {
			fmt.Printf("U+%04X\t%c\t%s\n", uchar, uchar, name)
		}

	}
}
