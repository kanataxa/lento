package parsing

import (
	"testing"
)

func TestListParser_List(t *testing.T) {
	t.Run("simply sentences [ <elements> ]", func(t *testing.T) {
		sentences := "[ this, is,  simple, sentence]"
		parser := NewListParser(sentences)

		if err := parser.Parse(); err != nil {
			t.Error(err)
		}
	})
	t.Run("bad sentences [ <elements> ,]", func(t *testing.T) {
		sentences := "[ this, is,  simple, sentence,]"
		parser := NewListParser(sentences)

		if err := parser.Parse(); err == nil {
			t.Errorf("this test case is never passed. bad err is %v", err)
		}
	})
}
