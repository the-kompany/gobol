package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// var (
// 	Token []string
// )

type itemType int

type Token struct {
	typ itemType
	val string
}

const (
	itemError itemType = iota
	itemDot
	itemEOF
	itemNumber
	itemPrint //
	itemIF
	itemElse
	itemEnd
	itemField  //identifier
	itemString // quoted string
	itemText   //plain text
	itemKey
	itemEqual
	itemValue
)

const eof rune = -1

//lexer holds the state of the scanner
type lexer struct {
	name  string //used only for error reports
	input string //the string being scanned
	start int    //start position of the item
	pos   int    //current position in the item
	width int
	items chan Token //channel of scanned item
	line  int
}

//stateFn represents the state of the scanner
//as a function that returns the next state
type stateFn func(*lexer) stateFn

func main() {

	//TODO get the filename from argument

	f, err := os.Open("hello.gbl")

	if err != nil {
		log.Fatal(err)
	}
	var extension = filepath.Ext("hello.gbl")

	if extension != ".gbl" {
		fmt.Printf("%s", "not a valid gobol file")
	}

	//temporary for testing > use scanner instead?
	// b := make([]byte, 1024)
	// f.Read(b)

	// tkn := strings.Fields(string(b))

	// for _, v := range tkn {
	// 	// if strings.HasPrefix(v, "\"") {
	// 	fmt.Println(v)
	// 	// }
	// }
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		fmt.Println(scanner.Text()) // token in unicode-char

		l := lex("", scanner.Text())

		fmt.Println(l.items)
	}

	time.Sleep(time.Second * 2)
}

func lex(name, input string) *lexer {
	l := &lexer{
		name:  name,
		input: input,
		items: make(chan Token),
		line:  1,
	}

	go l.run() //concurrently run state machine
	return l
}

func (l *lexer) run() {

	for state := lexText; state != nil; {
		state = state(l)
	}

	close(l.items)
}

func (l *lexer) emit(t itemType) {

	log.Println("debug emit")
	l.items <- Token{t, l.input[l.start:l.pos]}
	l.start = l.pos

}

//lexText
func lexText(l *lexer) stateFn {

	return lexInsideAction
}

//
func (l *lexer) lexQuote() stateFn {

Loop:
	for {
		switch r := l.next(); {
		case r == eof && !strings.HasPrefix(string(r), "\""):

			return l.errorf("String value not closed properly")

		case strings.HasPrefix(string(r), "\""):
			break Loop

		}

	}

	l.emit(itemString)
	return lexInsideAction
}

func (l *lexer) next() rune {

	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}

	r, w := utf8.DecodeLastRuneInString(l.input[l.pos:])

	log.Fatal(l.input[l.pos:])
	l.width = w
	l.pos += l.width
	return r
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func lexInsideAction(l *lexer) stateFn {

	r := l.next()

	fmt.Println(l.input)
	fmt.Println(strconv.QuoteRune(r))

	// for {
	// 	switch r := l.next(); {

	// 	case r == eof || r == '\n':
	// 		return l.errorf("unclosed actionm")
	// 	case isSpace(r):
	// 		l.ignore()
	// 	case string(r) == "print":

	// 		l.emit(itemField)
	// 	default:
	// 		fmt.Println(string(r))
	// 	}
	// }
	return lexInsideAction
}

func isSpace(r rune) bool {
	if r == ' ' || r == '\t' {
		return true
	}

	return false
}

func (l *lexer) backup() {
	l.pos = l.width
}

func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (i Token) String() string {

	fmt.Println("string method called")

	switch i.typ {
	case itemEOF:
		return "EOF"
	case itemError:
		return i.val

	}

	if len(i.val) > 10 {
		return fmt.Sprintf("%  10q...", i.val)
	}

	return fmt.Sprintf("%q", i.val)

}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- Token{itemError, fmt.Sprintf(format, args...)}
	return nil
}
