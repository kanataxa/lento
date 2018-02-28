package parsing

import (
	"fmt"
	"testing"
)

func TestListLexer_Parse(t *testing.T) {
	listLexer := NewListLexer("  [] a] bd] ,")
	listLexer.Next()
	tok := listLexer.Token()
	for tok.TokType != eof {
		fmt.Println(tok)
		listLexer.Next()
		tok = listLexer.Token()
	}
	fmt.Println(tok)
}
