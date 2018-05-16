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
 <element> ::= ('a'|...|'Z')[('a'|...|'Z')]*|<element>'='<element>|<list>;
*/
type MultiParser struct {
	Lexer
	toks []*Token
	size uint32
	pos  uint32
}

func NewMultiParser(input string, size uint32) Parser {
	lexer := NewListLexer(input)
	toks := make([]*Token, size)
	for i := 0; i < int(size); i++ {
		if i > lexer.Len() {
			break
		}
		toks[i] = lexer.Token()
		lexer.Next()
	}

	return &MultiParser{
		Lexer: lexer,
		toks:  toks,
		size:  size,
	}
}

func (p *MultiParser) consume() {
	p.toks[p.pos] = p.Token()
	p.Next()
	p.pos = p.nextPos()
}

func (p *MultiParser) nextPos() uint32 {
	return (p.pos + 1) % p.size
}
func (p *MultiParser) Match(tokType int) error {
	if tokType != p.toks[p.pos].TokType {
		return errors.Errorf("unknown token: [%d]", tokType)
	}
	p.consume()
	return nil
}

func (p *MultiParser) Parse() error {
	return errors.Wrap(p.list(), "failed parse")
}

func (p *MultiParser) list() error {
	if err := p.Match(lbrack); err != nil {
		return err
	}
	if err := p.elements(); err != nil {
		return errors.Wrap(err, "failed parse list")
	}
	return p.Match(rbrack)
}

func (p *MultiParser) elements() error {
	if err := p.element(); err != nil {
		return errors.Wrap(err, "failed parse elements")
	}
	for p.toks[p.pos].TokType == comma {
		if err := p.Match(comma); err != nil {
			return err
		}
		if err := p.element(); err != nil {
			return errors.Wrap(err, "failed parse elements")
		}
	}
	return nil
}

func (p *MultiParser) element() error {
	tokType := p.toks[p.pos].TokType
	if tokType == lbrack {
		return errors.Wrap(p.list(), "failed parse element")
	}
	if tokType == name && p.toks[p.nextPos()].TokType == equal {
		if err := p.Match(name); err != nil {
			return errors.Wrap(err, "failed parse element")
		}
		if err := p.Match(equal); err != nil {
			return errors.Wrap(err, "failed parse element")
		}
		if err := p.Match(name); err != nil {
			return errors.Wrap(err, "failed parse element")
		}

		return nil
	}
	if tokType == name {
		return errors.Wrap(p.Match(name), "failed parse element")
	}
	return errors.Errorf("unknown tokType [%d]", tokType)
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
