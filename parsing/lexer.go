package parsing

import (
	"fmt"
	"unicode/utf8"
)

const (
	na = iota
	eof
	name
	comma
	lbrack
	rbrack
)

const (
	EOFText = "-1"
)

type Lexer interface {
	Len() int
	Next()
	Token() *Token
}

type ListLexer struct {
	tokNames []string
	Input    string
	Pos      int
	PosText  string
	tok      *Token
}

func NewListLexer(input string) Lexer {
	tokType := na
	posText := input[0:1]
	if posText == "[" {
		tokType = lbrack
	} else if posText == "]" {
		tokType = rbrack
	} else if posText == "," {
		tokType = comma
	}
	return &ListLexer{
		tokNames: []string{"n/a", "<EOF>", "NAME", "COMMA", "LBRACK", "RBLACK"},
		Input:    input,
		Pos:      0,
		tok:      NewToken(tokType, posText),
		PosText:  input[1:2],
	}
}
func (l *ListLexer) Len() int {
	return utf8.RuneCountInString(l.Input)
}

func (l *ListLexer) Next() {
	for l.PosText != EOFText {
		switch l.PosText {
		case " ", "\t", "\n", "\r":
			l.WhiteSpace()
			continue
		case ",":
			l.tok = NewToken(comma, ",")
			l.Consume()
			return
		case "[":
			l.tok = NewToken(lbrack, "[")
			l.Consume()
			return
		case "]":
			l.tok = NewToken(rbrack, "]")
			l.Consume()
			return
		default:
			if !l.IsLetter() {
				panic(fmt.Errorf("unknown character: %s", l.PosText))
			}
			l.tok = NewToken(name, l.LetterName())
			return
		}
	}
	l.tok = NewToken(eof, "EOF")
}

func (l *ListLexer) Token() *Token {
	return l.tok
}

func (l *ListLexer) WhiteSpace() {
	for l.PosText == " " || l.PosText == "\t" || l.PosText == "\n" || l.PosText == "\r" {
		l.Consume()
	}
}
func (l *ListLexer) Consume() {
	l.Pos++
	if l.Pos >= utf8.RuneCountInString(l.Input) {
		l.PosText = EOFText
	} else {
		l.PosText = l.Input[l.Pos : l.Pos+1]
	}
}

func (l *ListLexer) IsLetter() bool {
	if l.PosText >= "a" && l.PosText <= "z" {
		return true
	}
	return l.PosText >= "A" && l.PosText <= "Z"
}

func (l *ListLexer) LetterName() string {
	var letterName string
	for l.IsLetter() {
		letterName += l.PosText
		l.Consume()
	}
	return letterName
}
