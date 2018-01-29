package parser

import (
	"fmt"
	"testing"
)

func TestRecursiveDecentParser_Number(t *testing.T) {
	t.Run("11+24+32-68", func(t *testing.T) {
		p := NewRecursiveDecentParser(NewSource("11 + 24 +      32 - 68"))
		output := p.Expr()
		if output != -1 {
			t.Error("output number not equal")
			t.Log("expect:", -1, "actual:", output)
		}

		fmt.Println("succeed test, result: ", output)
	})
	t.Run("-1", func(t *testing.T) {
		p := NewRecursiveDecentParser(NewSource("-1"))
		output := p.Expr()
		if output != -1 {
			t.Error("output number not equal")
			t.Log("expect:", -1, "actual:", output)
		}

		fmt.Println("succeed test, result: ", output)
	})
	t.Run("10+20+20", func(t *testing.T) {
		p := NewRecursiveDecentParser(NewSource("-10+20+20"))
		output := p.Expr()
		if output != 30 {
			t.Error("output number not equal")
			t.Log("expect:", 30, "actual:", output)
		}

		fmt.Println("succeed test, result: ", output)
	})

	t.Run("10/10+20", func(t *testing.T) {
		p := NewRecursiveDecentParser(NewSource("10/10+20"))
		output := p.Expr()
		if output != 21 {
			t.Error("output number not equal")
			t.Log("expect:", 21, "actual:", output)
		}

		fmt.Println("succeed test, result: ", output)
	})

	t.Run("10*10/20", func(t *testing.T) {
		p := NewRecursiveDecentParser(NewSource("10*10/ 20"))
		output := p.Expr()
		if output != 5 {
			t.Error("output number not equal")
			t.Log("expect:", 5, "actual:", output)
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
	t.Run("(10+10)*20", func(t *testing.T) {
		p := NewRecursiveDecentParser(NewSource("(10+10)/20"))
		output := p.Expr()
		if output != 1 {
			t.Error("output number not equal")
			t.Log("expect:", 1, "actual:", output)
		}

		fmt.Println("succeed test, result: ", output)
	})
	t.Run("(10+10*40*(10+10))/10", func(t *testing.T) {
		p := NewRecursiveDecentParser(NewSource("(10+10*40*(10+10))/10"))
		output := p.Expr()
		if output != 801 {
			t.Error("output number not equal")
			t.Log("expect:", 801, "actual:", output)
		}

		fmt.Println("succeed test, result: ", output)
	})

}
