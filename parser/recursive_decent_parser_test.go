package parser

import (
	"fmt"
	"testing"
)

func TestRecursiveDecentParser_Number(t *testing.T) {
	t.Run("11+24+32-68", func(t *testing.T) {
		p := NewRecursiveDecentParser(NewSource("11+24+32-68"))
		output := p.Expr()
		if output != -1 {
			t.Error("output number not equal")
			t.Log("expect:", -1, "actual:", output)
		}

		fmt.Println("succeed test, result: ", output)
	})
	t.Run("10+20+20", func(t *testing.T) {
		p := NewRecursiveDecentParser(NewSource("10+20+20"))
		output := p.Expr()
		if output != 50 {
			t.Error("output number not equal")
			t.Log("expect:", 50, "actual:", output)
		}

		fmt.Println("succeed test, result: ", output)
	})
	t.Run("10*20+20", func(t *testing.T) {
		p := NewRecursiveDecentParser(NewSource("10*20+20"))
		output := p.Expr()
		if output != 220 {
			t.Error("output number not equal")
			t.Log("expect:", 220, "actual:", output)
		}

		fmt.Println("succeed test, result: ", output)
	})
	t.Run("10*10*20", func(t *testing.T) {
		p := NewRecursiveDecentParser(NewSource("10*10*20"))
		output := p.Expr()
		if output != 2000 {
			t.Error("output number not equal")
			t.Log("expect:", 2000, "actual:", output)
		}

		fmt.Println("succeed test, result: ", output)
	})
	t.Run("10+10*20", func(t *testing.T) {
		p := NewRecursiveDecentParser(NewSource("10+10*20"))
		output := p.Expr()
		if output != 210 {
			t.Error("output number not equal")
			t.Log("expect:", 210, "actual:", output)
		}

		fmt.Println("succeed test, result: ", output)
	})
}
