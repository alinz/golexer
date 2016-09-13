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

	lexer := New(file, 1024)

	for i := 0; i < 100; i++ {
		lexer.AcceptRunUntil("\n")
		fmt.Printf("->%s<- %d\n", lexer.CurrentString(), len(lexer.CurrentString()))
		lexer.Ignore()

		lexer.Next()
		lexer.Ignore()
	}

	// lexer.AcceptRunUntil(" ")
	// fmt.Println(lexer.CurrentString())
	// lexer.Ignore()

	// lexer.AcceptRunUntil(" ")
	// fmt.Println(lexer.CurrentString())
	// lexer.Ignore()
}
