package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Data struct {
	Vars map[string]string
	Line int
}

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
	d := &Data{}
	d.Vars = make(map[string]string)

	for scanner.Scan() {

		lower := scanner.Text()

		if strings.HasPrefix(lower, "move") || strings.HasPrefix(lower, "MOVE") {
			d.handleMove(lower)
			continue
		}

		if strings.HasPrefix(lower, "display") || strings.HasPrefix(lower, "DISPLAY") {
			d.handleDisplay(lower)
		}
	}

}

func (d *Data) handleMove(val string) {

	trimmed := strings.TrimSpace(val)

	splitted := split(trimmed)

	toLow := strings.ToLower(splitted[2])
	if toLow != "to" || len(splitted) < 4 {
		fmt.Println("Error: Inavalid syntax for MOVE")
		os.Exit(1)
	}

	if strings.HasPrefix(splitted[1], "\"") {
		d.Vars[splitted[3]] = splitted[1]
	} else {
		if _, ok := d.Vars[splitted[1]]; !ok {
			fmt.Println("Error: Undefined variable \"", splitted[1], "\"")
			os.Exit(1)
		}

	}

}

func (d *Data) handleDisplay(val string) {
	trimmed := strings.TrimSpace(val)

	splitted := split(trimmed)

	if strings.HasPrefix(splitted[1], "\"") {

		trimmedQuote := splitted[1][1 : len(splitted[1])-1]
		fmt.Println(trimmedQuote)

	} else {
		if _, ok := d.Vars[splitted[1]]; !ok {
			fmt.Println("Error: Undefined variable \"", splitted[1], "\"")
			os.Exit(1)
		}

		varData := d.Vars[splitted[1]]
		trimmedQuote := varData[1 : len(varData)-1]
		fmt.Println(trimmedQuote)

	}
}

func split(val string) []string {
	// pos := 0
	splitted := strings.Split(val, " ")

	fields := []string{}

	var ok = true
	var s string

	for _, v := range splitted {
		if strings.HasPrefix(v, "\"") {
			s += v

			ok = false
			continue
		}

		if !ok {

			s += " "
			s += v

			if strings.HasSuffix(v, "\"") {
				ok = true
				fields = append(fields, s)
				continue
			}
		}

		fields = append(fields, v)

	}

	return fields

}
