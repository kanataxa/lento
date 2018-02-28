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

type Lexer struct {
	Input   string
	Pos     int
	PosText string
	EOF     string
	EOFType int
}

type ListLexer struct {
	name     int
	comma    int
	lbrack   int
	rbrack   int
	tokNames []string
	*Lexer
}

func NewListLexer(input string) *ListLexer {
	return &ListLexer{
		tokNames: []string{"n/a", "<EOF>", "NAME", "COMMA", "LBRACK", "RBLACK"},
		Lexer: &Lexer{
			Input:   input,
			EOF:     "-1",
			EOFType: 1,
			PosText: input[0:1],
		},
	}
}
func (ll *ListLexer) TokenName(x int) string {
	return ll.tokNames[x]
}

func (l *Lexer) NextToken() *Token {
	for l.PosText != l.EOF {
		switch l.PosText {
		case " ", "\t", "\n", "\r":
			l.WhiteSpace()
			continue
		case ",":
			l.Consume()
			return NewToken(comma, ",")
		case "[":
			l.Consume()
			return NewToken(lbrack, "[")
		case "]":
			l.Consume()
			return NewToken(rbrack, "]")
		default:
			if l.IsLetter() {
				return NewToken(name, l.LetterName())
			}
			panic(fmt.Errorf("unknown character: %s", l.PosText))
		}
	}
	return NewToken(l.EOFType, "EOF")
}

func (l *Lexer) WhiteSpace() {
	for l.PosText == " " || l.PosText == "\t" || l.PosText == "\n" || l.PosText == "\r" {
		l.Consume()
	}
}
func (l *Lexer) Consume() {
	l.Pos++
	if l.Pos >= utf8.RuneCountInString(l.Input) {
		l.PosText = l.EOF
	} else {
		l.PosText = l.Input[l.Pos : l.Pos+1]
	}
}

func (l *Lexer) IsLetter() bool {
	if l.PosText >= "a" && l.PosText <= "z" {
		return true
	}
	return l.PosText >= "A" && l.PosText <= "Z"
}

func (l *Lexer) LetterName() string {
	var letterName string
	for l.IsLetter() {
		letterName += l.PosText
		l.Consume()
	}
	return letterName
}
