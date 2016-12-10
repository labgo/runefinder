package main

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	var tests = []struct {
		input string
		want  []string
	}{
		{"AMPERSAND", []string{"AMPERSAND"}},
		{"PLUS SIGN", []string{"PLUS", "SIGN"}},
		{"HYPHEN-MINUS", []string{"HYPHEN", "MINUS"}},
	}
	for _, test := range tests {
		if got := tokenize(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("tokenize(%q) = %v", test.input, got)
		}
	}
}

// 0026;AMPERSAND;Po;0;ON;;;;;N;;;;;
func TestParse(t *testing.T) {

	var tests = []struct {
		input     string
		wantChar  rune
		wantWords []string
	}{
		{"0026;AMPERSAND;Po;0;ON;;;;;N;;;;;", '&', []string{"AMPERSAND"}},
		{"0021;EXCLAMATION MARK;Po;0;ON;;;;;N;;;;;", '!', []string{"EXCLAMATION", "MARK"}},
		{"002E;FULL STOP;Po;0;CS;;;;;N;PERIOD;;;;", '.', []string{"FULL", "STOP", "PERIOD"}},
		{"0027;APOSTROPHE;Po;0;ON;;;;;N;APOSTROPHE-QUOTE;;;", '\'', []string{"APOSTROPHE", "QUOTE"}},
	}
	for _, test := range tests {
		if gotChar, gotWords := Parse(test.input); gotChar != test.wantChar ||
			!reflect.DeepEqual(gotWords, test.wantWords) {
			t.Errorf("Parse(%q) = (%v, %v)", test.input, gotChar, gotWords)
		}
	}
}
