package parser

import (
	"strconv"
	"unicode"
)

const (
	plus  = "+"
	minus = "-"
	muluti = "*"
	division = "/"
)

type RecursiveDecentParser struct {
	source *Source
}

// TODO: use interface
func NewRecursiveDecentParser(s *Source) *RecursiveDecentParser {
	return &RecursiveDecentParser{
		source: s,
	}
}


// <expr> ::= <term> [ ('+' | '-') <term> ]*s
// <term> ::= <factor> [ ('*' | '/') <factor> ]*
// <factor> ::= <number> | '(' <expr> ')'
// <number> ::= ('1'|...|'9')[('0'|...|'9')]*
func (p *RecursiveDecentParser) Expr() int {
	x := p.term()
	for {
		if p.source.Pos() >= p.source.Len() {
			break
		}
		switch p.source.PosText() {
		case plus:
			p.source.Next()
			x += p.term()
		case minus:
			p.source.Next()
			x -= p.term()
		}
	}
	return x
}

func (p *RecursiveDecentParser) term() int {
	x := p.factor()
	for {
		if p.source.Pos() >= p.source.Len() {
			break
		}
		switch p.source.PosText() {
		case muluti:
			p.source.Next()
			x *= p.factor()
			continue
		case division:
			p.source.Next()
			x /= p.factor()
			continue
		default:
		}
		break
	}
	return x
}

func (p *RecursiveDecentParser) factor() int {
	x := p.number()
	return x
}

func (p *RecursiveDecentParser) number() int {
	var text string
	source := p.source.Text()
	for _, r := range source {
		if !unicode.IsNumber(r) {
			break
		}
		text += string(r)
		p.source.Next()
	}
	num, _ := strconv.Atoi(text)
	return num
}