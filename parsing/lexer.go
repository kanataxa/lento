package parsing

import (
	"fmt"
	"unicode/utf8"
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
		name:     2,
		comma:    3,
		lbrack:   4,
		rbrack:   5,
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

func (ll *ListLexer) NextToken() *Token {
	for ll.PosText != ll.EOF {
		switch ll.PosText {
		case " ", "\t", "\n", "\r":
			ll.WhiteSpace()
			continue
		case ",":
			ll.Consume()
			return NewToken(ll.comma, ",")
		case "[":
			ll.Consume()
			return NewToken(ll.lbrack, "[")
		case "]":
			ll.Consume()
			return NewToken(ll.rbrack, "]")
		default:
			if ll.IsLetter() {
				return NewToken(ll.name, ll.LetterName())
			}
			panic(fmt.Errorf("Unknown character: %s", ll.PosText))
		}
	}
	return NewToken(ll.EOFType, "EOF")
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
