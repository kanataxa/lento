package parsing

import (
	"fmt"
	"testing"
)

func TestListLexer_Parse(t *testing.T) {
	listLexer := NewListLexer("  [] a] bd] ,")
	tok := listLexer.NextToken()
	for tok.TokType != listLexer.EOFType {
		fmt.Println(tok)
		tok = listLexer.NextToken()
	}
	fmt.Println(tok)
}
