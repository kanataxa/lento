package parsing

type Token struct {
	TokType int
	Text    string
}

func NewToken(tokType int, text string) *Token {
	return &Token{
		TokType: tokType,
		Text:    text,
	}
}

func (t *Token) String() string {
	return "<" + t.Text + ">"
}
