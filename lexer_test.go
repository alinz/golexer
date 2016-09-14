package golexer

import (
	"bytes"
	"io"
	"testing"
	"unicode/utf8"
)

func stream(str string) (io.Reader, []rune) {
	var runes []rune

	for i, w := 0, 0; i < len(str); i += w {
		runeValue, width := utf8.DecodeRuneInString(str[i:])
		runes = append(runes, runeValue)
		w = width
	}

	input := bytes.NewBufferString(str)

	return input, runes
}

func TestLexerNext(t *testing.T) {
	input, runes := stream(`Hello World سلام دنیا こんにちは世界`)

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

func TestIgnore(t *testing.T) {

}
