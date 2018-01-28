package parser

import (
	"strconv"
	"unicode"
)

const (
	plus  = "+"
	minus = "-"
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

func (p *RecursiveDecentParser) Expr() int {
	x := p.number()
	for {
		if p.source.Pos() >= p.source.Len() {
			break
		}
		switch p.source.PosText() {
		case plus:
			p.source.Next()
			x += p.number()
		case minus:
			p.source.Next()
			x -= p.number()
		}
	}
	return x
}
