package main

import (
	"testing"

	"gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { check.TestingT(t) }

type MySuite struct{}

var _ = check.Suite(&MySuite{})

func (s *MySuite) TestFindOneWord(c *check.C) {
	index := map[string][]rune{
		"REGISTERED": []rune{0xAE},
	}

	tests := map[string][]rune{
		"registered": []rune{0xAE},
		"nonesuch":   []rune{},
	}
	for query, found := range tests {
		c.Assert(findRunes(query, index), check.DeepEquals, found)
	}
}

func (s *MySuite) TestIndexHyphenatedWord(c *check.C) {
	index, _ := indexWords("002D;HYPHEN-MINUS;Pd;0;ES;;;;;N;;;;;")

	tests := map[string][]rune{
		"hyphen": []rune{0x2D},
	}
	for query, found := range tests {
		c.Assert(findRunes(query, index), check.DeepEquals, found)
	}
}

func (s *MySuite) TestIndexOldNameField(c *check.C) {
	index, _ := indexWords(
		"0028;LEFT PARENTHESIS;Ps;0;ON;;;;;Y;OPENING PARENTHESIS;;;;")

	tests := map[string][]rune{
		"opening": []rune{0x28},
	}
	for query, found := range tests {
		c.Assert(findRunes(query, index), check.DeepEquals, found)
	}
}
