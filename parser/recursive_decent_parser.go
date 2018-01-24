package parser

import (
	"unicode"
)

const (
	plusExpr  = "+"
	minusExpr = "-"
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

func (p *RecursiveDecentParser) Number() string {
	source := p.source.Text(0, p.source.Len())
	for _, str := range source {
		if !unicode.IsNumber(rune(str)) {
			return p.source.Text(0, p.source.Pos())
		}
		p.source.Next()
	}
	return source
}
