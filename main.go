package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/the-kompany/gobol/transpiler"
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
	d := &transpiler.Data{}
	d.Vars = make(map[string]string)

	//scan eeach line and append to slice
	//better for parsing
	for scanner.Scan() {

		l := scanner.Text()
		trimmed := strings.TrimSpace(l)

		if strings.HasPrefix(trimmed, "//") {
			continue
		} else if strings.Contains(trimmed, "//") {
			i := strings.Index(trimmed, "//")
			if trimmed[i-1] != '"' {
				splitted := strings.Split(trimmed, "//")
				d.Lines = append(d.Lines, strings.TrimSpace(splitted[0]))

			}
		} else if len(trimmed) < 1 {
			continue
		}

		d.Lines = append(d.Lines, trimmed)

	}

	// log.Println(len(d.Lines))

	//TODO Parse it

	for _, v := range d.Lines {
		if strings.HasPrefix(v, "move") || strings.HasPrefix(v, "MOVE") {
			d.HandleMove(v)
			continue
		}

		if strings.HasPrefix(v, "display") || strings.HasPrefix(v, "DISPLAY") {
			d.HandleDisplay(v)
		}

		if strings.HasPrefix(v, "upshift") || strings.HasPrefix(v, "UPSHIFT") {
			d.HandleUpShift(v, "", "")
		}
	}

}
