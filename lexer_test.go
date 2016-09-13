package golexer

import (
	"bytes"
	"testing"
	"unicode/utf8"
)

func TestLexerNext(t *testing.T) {
	var runes []rune

	const str = `Hello World سلام دنیا`
	for i, w := 0, 0; i < len(str); i += w {
		runeValue, width := utf8.DecodeRuneInString(str[i:])
		runes = append(runes, runeValue)
		w = width
	}

	input := bytes.NewBufferString(str)

	lexer := New(input, 10)

	var val rune
	for i := 0; i < len(runes); i++ {
		val = lexer.Next()
		if runes[i] != val {
			t.Errorf("got '%s' instead of '%s", string(runes[i]), string(val))
		}
	}

	val = lexer.Next()
	if val != 0 {
		t.Errorf("it should be 0 but got '%s'", string(val))
	}
}
