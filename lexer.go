package golexer

import (
	"io"
	"strings"
	"unicode/utf8"
)

type Lexer struct {
	input       io.Reader
	buffer      []byte
	fixedBuffer []byte
	pos         int
	width       int
	closed      bool
}

// Ignore resets buffer upto pos pointer, resets all internal vars
func (l *Lexer) Ignore() {
	l.buffer = l.buffer[l.pos:]
	l.pos = 0
	l.width = 0
}

// Backup moves the position one step backward
func (l *Lexer) Backup() {
	l.pos -= l.width
	l.width = 0
	if l.pos < 0 {
		l.pos = 0
	}
}

// Current return the current value that pointer points to without advancing the pointer
func (l *Lexer) Current() rune {
	l.Backup()
	return l.Next()
}

// CurrentString returns the current representation of buffer in string
func (l *Lexer) CurrentString() string {
	return string(l.buffer[:l.pos])
}

// Peek looks ahaed one byte without advancing position
func (l *Lexer) Peek() rune {
	val := l.Next()
	l.Backup()
	return val
}

// PeekNth looks ahead of nth rune in input, without advancing position
// returns zero if n is zero or n is bigger than input
func (l *Lexer) PeekNth(n int) rune {
	var val rune

	// keeps pos and width of current location
	// because we need to set it back
	pos := l.pos
	width := l.width

	if n <= 0 {
		return val
	}

	for i := 0; i < n; i++ {
		val = l.Next()
		// if val is 0, it means that we are reached the end of the input.
		// break the loop
		if val == 0 {
			break
		}
	}

	l.pos = pos
	l.width = width

	return val
}

// AcceptRunUntil it does exatcly as AcceptRun but in reverse way. It consumes runes
// until it find one that match the notValid one.
// as an example you want to read everythign until '\n'
func (l *Lexer) AcceptRunUntil(notValid string) {
	val := l.Next()
	for val != 0 && strings.IndexRune(notValid, val) == -1 {
		val = l.Next()
	}
	//we only going back if one of the chars found. if we get 0
	//it means that some kind of error happens, either Reader closed
	//or we reached end of the input. so we do not want to go back
	if val != 0 {
		l.Backup()
	}
}

func (l *Lexer) Next() rune {
	length := len(l.buffer)

	if length < l.pos+4 && !l.closed {
		l.read()
		return l.Next()
	}

	if l.pos >= length && l.closed {
		return 0
	}

	val, size := utf8.DecodeRune(l.buffer[l.pos:])
	l.width = size
	l.pos += size

	return val
}

func (l *Lexer) read() {
	//we don not need to reset the l.fixedBuffer
	n, err := l.input.Read(l.fixedBuffer)

	//adding fixedBuffer to buffer
	for i := 0; i < n; i++ {
		l.buffer = append(l.buffer, l.fixedBuffer[i])
	}

	if err != nil {
		l.closed = true
		return
	}
}

func New(input io.Reader) *Lexer {
	return &Lexer{
		input:       input,
		buffer:      nil,
		fixedBuffer: make([]byte, 4),
		closed:      false,
	}
}
