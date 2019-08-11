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
	var trimmed string

	//scan eeach line and append to slice
	//better for parsing
	for scanner.Scan() {

		l := scanner.Text()
		if ifStart == false {
			trimmed = strings.TrimSpace(l)
		}

		if strings.HasPrefix(trimmed, "//") {
			continue
		} else if strings.Contains(trimmed, "//") {

			i := strings.Index(trimmed, "//")
			if trimmed[i-1] != '"' {
				splitted := strings.Split(trimmed, "//")
				d.Lines = append(d.Lines, strings.TrimSpace(splitted[0]))
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
				d.Lines = append(d.Lines, ifBlock)
				continue
			}

			ifBlock += l

			continue
		}

		d.Lines = append(d.Lines, trimmed)

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

		token := strings.SplitAfter(v, " ")

		tokenTrimmed := strings.ToLower(strings.TrimSpace(token[0]))

		var firstToken string
		if strings.Contains(tokenTrimmed, "(") {
			spl := strings.Split(tokenTrimmed, "(")
			firstToken = spl[0]
		} else {
			firstToken = tokenTrimmed
		}

		switch firstToken {
		case "if":
			if !parser.ValidIfBlock(v) {
				fmt.Println("Syntax error: invalid if block at line ")
				os.Exit(1)
			}

			d.IfBlock(v)

		case "move":
			d.Move(v)
			continue
		case "upshift":
			d.Shift(v, "", 1)
		case "downshift":
			d.Shift(v, "", 0)
		case "display":
			d.Display(v)
		default:
			fmt.Printf("Error: Undefined %v in %v", firstToken, v)
		}
	}

}
