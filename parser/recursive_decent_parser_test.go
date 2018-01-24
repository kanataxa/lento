package parser

import (
	"fmt"
	"testing"
)

func TestRecursiveDecentParser_Number(t *testing.T) {
	p := NewRecursiveDecentParser(NewSource("11+24+32"))
	fmt.Println(p.Number())
}
