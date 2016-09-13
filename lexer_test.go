package golexer

import (
	"fmt"
	"os"
	"testing"
)

func TestLexer(t *testing.T) {
	file, err := os.Open("./sample/test.txt")
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	lexer := New(file)

	lexer.AcceptRunUntil(" ")
	fmt.Println(lexer.CurrentString())
	lexer.Ignore()

	lexer.Next()
	lexer.Ignore()

	p := lexer.PeekNth(6)
	fmt.Println(string(p))

	lexer.AcceptRunUntil(" ")
	fmt.Println(lexer.CurrentString())
	lexer.Ignore()

	lexer.Next()
	lexer.Ignore()

	lexer.AcceptRunUntil(" ")
	fmt.Println(lexer.CurrentString())
	lexer.Ignore()

	lexer.Next()
	lexer.Ignore()

	lexer.AcceptRunUntil(" ")
	fmt.Println(lexer.CurrentString())
	lexer.Ignore()

	lexer.Next()
	lexer.Ignore()

	// lexer.AcceptRunUntil(" ")
	// fmt.Println(lexer.CurrentString())
	// lexer.Ignore()

	// lexer.AcceptRunUntil(" ")
	// fmt.Println(lexer.CurrentString())
	// lexer.Ignore()
}
