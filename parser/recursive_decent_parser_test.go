package parser

import (
	"fmt"
	"testing"
)

func TestRecursiveDecentParser_Number(t *testing.T) {
	p := NewRecursiveDecentParser(NewSource("11+24+32-68"))
	output := p.Expr()
	if output != -1 {
		t.Error("output number not equal")
		t.Log("expect:", -1, "actual:", output)
	}

	fmt.Println("succeed test, result: ", output)
}
