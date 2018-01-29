package parser

import (
	"fmt"
	"strconv"
	"unicode"
)

const (
	plus     = "+"
	minus    = "-"
	multi    = "*"
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
// <factor> ::= [' ']*<unaryNumber> | '(' <expr> ')'[' ']*
// <unaryNumber> ::= ['+'|'-']<number>
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
			continue
		case minus:
			p.source.Next()
			x -= p.term()
			continue
		}
		break
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
		case multi:
			p.source.Next()
			x *= p.factor()
			continue
		case division:
			p.source.Next()
			x /= p.factor()
			continue
		}
		break
	}
	return x
}

func (p *RecursiveDecentParser) factor() int {
	var x int
	p.spaces()
	if p.source.PosText() == "(" {
		p.source.Next()
		x = p.Expr()
		p.source.Next()
	} else {
		x = p.unaryNumber()
	}
	p.spaces()
	return x
}

func (p *RecursiveDecentParser) spaces() {
	for p.source.Pos() < p.source.Len() && p.source.PosText() == " " {
		p.source.Next()
	}
}

func (p *RecursiveDecentParser) unaryNumber() int {
	switch p.source.PosText() {
	case plus:
		p.source.Next()
		return p.number()
	case minus:
		p.source.Next()
		return -p.number()
	}
	return p.number()
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
	num, err := strconv.Atoi(text)
	if err != nil {
		panic(fmt.Errorf(err.Error()+" text:%s, soure:%s", text, source))
	}
	return num
}
