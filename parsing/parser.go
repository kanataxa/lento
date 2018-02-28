package parsing

import (
	"github.com/pkg/errors"
)

type Parser interface {
	Lexer
	Parse() error
}

/*
 <list> ::= '['<elements>']';
 <elements> ::= <element>[','<element>]*;
 <element> ::= ('a'|...|'Z')[('a'|...|'Z')]*|<list>;
*/
type ListParser struct {
	laTok *Token
	Lexer
}

func (p *ListParser) Match(tokType int) error {
	if tokType != p.laTok.TokType {
		return errors.Errorf("unknown tokType %d", tokType)
	}
	p.consume()
	return nil
}

func (p *ListParser) consume() {
	p.Next()
	p.laTok = p.Token()
}

func (p *ListParser) Parse() error {
	return errors.Wrap(p.List(), "failed parse")
}

func (p *ListParser) List() error {
	if err := p.Match(lbrack); err != nil {
		return err
	}
	if err := p.Elements(); err != nil {
		return errors.Wrap(err, "failed parse list")
	}
	return p.Match(rbrack)
}

func (p *ListParser) Elements() error {
	if err := p.Element(); err != nil {
		return errors.Wrap(err, "failed parse elements")
	}
	for p.laTok.TokType == comma {
		if err := p.Match(comma); err != nil {
			return err
		}
		if err := p.Element(); err != nil {
			return errors.Wrap(err, "failed parse elements")
		}
	}
	return nil
}

func (p *ListParser) Element() error {
	if p.laTok.TokType == name {
		return p.Match(name)
	}
	if p.laTok.TokType == lbrack {
		return errors.Wrap(p.List(), "failed parse element")
	}
	return errors.Errorf("unknown tokType %s", p.laTok.TokType)
}

func NewListParser(input string) Parser {
	lexer := NewListLexer(input)
	return &ListParser{
		Lexer: lexer,
		laTok: lexer.Token(),
	}
}
