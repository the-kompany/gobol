package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/the-kompany/gobol/pkg/handler"
	"github.com/the-kompany/gobol/pkg/parser"
)

var itemType int

//For keyword/identifier
const (
	Move    = "move"
	UPSHIFT = "upshift"
	DISPLAY = "display"
)

func main() {

	//TODO get the filename from argument

	args := os.Args

	if len(args) < 2 {
		fmt.Println("No file name provided in argument")
		os.Exit(1)
	}

	fileName := os.Args[1]

	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	var extension = filepath.Ext("hello.gbl")

	if extension != ".gbl" {
		fmt.Printf("%s", "not a valid gobol file")
	}

	scanner := bufio.NewScanner(f)
	d := &handler.Data{}
	d.Vars = make(map[string]string)

	var ifBlock string
	ifStart := false
	var performstart bool
	var performBlock string
	var trimmed string
	var lineNumber int

	var recordBlock string
	var recordStart bool

	//scan eeach line and append to slice
	//better for parsing
	for scanner.Scan() {

		l := scanner.Text()
		lineNumber++
		if ifStart == false {
			trimmed = strings.TrimSpace(l)
		} else if performstart == false {
			trimmed = strings.TrimSpace(l)

		}

		if strings.HasPrefix(trimmed, "//") {
			continue
		} else if strings.Contains(trimmed, "//") {

			i := strings.Index(trimmed, "//")
			if trimmed[i-1] != '"' {
				splitted := strings.Split(trimmed, "//")
				d.Lines = append(d.Lines, handler.Token{Value: strings.TrimSpace(splitted[0]), Line: lineNumber})
				continue

			}
			//ignore the empty line
		} else if len(trimmed) < 1 {
			continue
		}

		// check if the string contains IF condition block
		if strings.HasPrefix(strings.ToLower(l), "if") {
			ifStart = true
			ifBlock += l
			continue
		}

		if ifStart == true {
			if strings.HasPrefix(strings.ToLower(l), "end-if") {
				ifStart = false
				ifBlock += l
				d.Lines = append(d.Lines, handler.Token{Value: ifBlock, Line: lineNumber})
				continue
			}

			ifBlock += l

			continue
		}

		if strings.HasPrefix(strings.ToLower(l), "perform") {
			performstart = true
			performBlock += l
			continue
		}

		if performstart == true {

			if strings.HasPrefix(strings.ToLower(l), "end-perform") {
				performstart = false
				performBlock += " " + l
				d.Lines = append(d.Lines, handler.Token{Value: performBlock, Line: lineNumber})
				performBlock = ""
				continue
			}

			performBlock += l

			continue
		}

		if strings.HasPrefix(strings.ToLower(l), "record") {
			recordStart = true
			recordBlock += l
			continue
		}

		if recordStart == true {
			if strings.HasPrefix(strings.ToLower(l), "end-record") {
				recordStart = false
				recordBlock += " " + l
				d.Lines = append(d.Lines, handler.Token{Value: recordBlock, Line: lineNumber})
				recordBlock = ""
				continue
			}

			recordBlock += l
			continue
		}

		d.Lines = append(d.Lines, handler.Token{Value: trimmed, Line: lineNumber})

	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, err)
			fmt.Println("eof")
		}
	}

	// log.Println(d.Lines[len(d.Lines)-1])

	//TODO Parse it

	for _, v := range d.Lines {

		token := strings.SplitAfter(v.Value, " ")

		tokenTrimmed := strings.ToLower(strings.TrimSpace(token[0]))

		var firstToken string
		if strings.Contains(tokenTrimmed, "(") {
			spl := strings.Split(tokenTrimmed, "(")
			firstToken = spl[0]
		} else {
			firstToken = tokenTrimmed
		}

		switch firstToken {
		case "perform":
			if tokens, err := parser.ValidPerformBlock(v.Value); err != nil {
				fmt.Println("Syntax error: invalid PERFORM block at line ", v.Line)
				os.Exit(1)
			} else {
				d.PerformLoopBlock(tokens)
			}

		case "if":
			if !parser.ValidIfBlock(v.Value) {
				fmt.Println("Syntax error: invalid if block at line ", v.Line)
				os.Exit(1)
			}

			d.IfBlock(v.Value)

		case "move":
			d.Move(v.Value)
			continue
		case "upshift":
			d.Shift(v.Value, "", 1)
		case "downshift":
			d.Shift(v.Value, "", 0)
		case "display":
			d.Display(v.Value)
		case "open":
			d.Open(v.Value)
		case "record":
			d.ExecuteRecord(v.Value)
		default:
			fmt.Printf("Error: Undefined %v at line %v \n", firstToken, v.Line)
		}
	}

}
