package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
type Pos int

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
	itemMove
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

		tkn := <-l.items

		fmt.Println(tkn)
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

	l.width = 0
	// l.pos = 3
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
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	// log.Println("next debug ", string(r))
	l.pos += l.width
	if r == '\n' {
		l.line++
	}
	return r
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func lexInsideAction(l *lexer) stateFn {

	// fmt.Println(strconv.QuoteRune(r))

	var keyWord string
	for {

		r := l.next()
		if isSpace(r) || r == '\n' || r == eof {
			break
		}
		keyWord += string(r)
	}

	switch keyWord {

	case "MOVE":
		log.Println("MOVE")
		l.emit(itemMove)

	default:
		return l.errorf("unrecognized character")
	}
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
