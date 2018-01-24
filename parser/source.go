package parser

import (
	"unicode/utf8"
)

type Source struct {
	text string
	pos  int
}

func (s *Source) Len() int {
	return utf8.RuneCountInString(s.text)
}

func (s *Source) Text(start, end int) string {
	return s.text[start:end]
}

func (s *Source) Pos() int {
	return s.pos
}

func (s *Source) Next() {
	s.pos++
}

func NewSource(text string) *Source {
	return &Source{
		text: text,
		pos:  0,
	}
}
