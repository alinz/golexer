package golexer

import (
	"bytes"
	"testing"
)

func TestLexerNext(t *testing.T) {
	input := bytes.NewBufferString(`Hello World`)

	matched := []rune{
		'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd',
	}

	lexer := New(input, 10)

	var val rune
	for i := 0; i < len(matched); i++ {
		val = lexer.Next()
		if matched[i] != val {
			t.Errorf("got '%s' instead of '%s", string(matched[i]), string(val))
		}
	}

	val = lexer.Next()
	if val != 0 {
		t.Errorf("it should be 0 but got '%s'", string(val))
	}
}
