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

func TestParse(t *testing.T) {

	var tests = []struct {
		input     string
		wantChar  rune
		wantName  string
		wantWords []string
	}{
		{"0026;AMPERSAND;Po;0;ON;;;;;N;;;;;",
			'&', "AMPERSAND", []string{"AMPERSAND"}},
		{"0021;EXCLAMATION MARK;Po;0;ON;;;;;N;;;;;",
			'!', "EXCLAMATION MARK", []string{"EXCLAMATION", "MARK"}},
		{"002E;FULL STOP;Po;0;CS;;;;;N;PERIOD;;;;",
			'.', "FULL STOP | PERIOD", []string{"FULL", "STOP", "PERIOD"}},
		{"0027;APOSTROPHE;Po;0;ON;;;;;N;APOSTROPHE-QUOTE;;;",
			'\'', "APOSTROPHE | APOSTROPHE-QUOTE", []string{"APOSTROPHE", "QUOTE"}},
	}
	for _, test := range tests {
		if gotChar, gotName, gotWords := Parse(test.input); gotChar != test.wantChar ||
				gotName != test.wantName ||
				!reflect.DeepEqual(gotWords, test.wantWords) {
			t.Errorf("Parse(%q) = (%v, %v, %v)", test.input, gotChar, gotName, gotWords)
		}
	}
}

func TestSliceHasAllStrings(t *testing.T) {
	var tests = []struct {
		inHaystack []string
		inNeedles  []string
		want  		 bool
	}{
		{[]string{"PLUS", "SIGN"}, []string{"SIGN", "PLUS"}, true},
		{[]string{"PLUS", "SIGN"}, []string{"SIGN"}, true},
		{[]string{"PLUS", "SIGN"}, []string{"SPAM"}, false},
		{[]string{"PLUS", "SIGN"}, []string{"PLUS", "SPAM"}, false},
	}
	for _, test := range tests {
		if got := sliceHasAllStrings(test.inHaystack, test.inNeedles); got != test.want {
			t.Errorf("siceHasAllStrings(%q, %q) = %v", test.inHaystack, test.inNeedles, got)
		}
	}
}
