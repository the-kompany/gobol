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

		if strings.HasPrefix(lower, "upshift") || strings.HasPrefix(lower, "UPSHIFT") {
			d.handleUpShift(lower, "", "")
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

func (d *Data) handleUpShift(val, first, to string) {
	//UPSHIFT(VAR)

	//get UPSHIFT function argument
	var (
		ok    = false
		start = false
		end   = false
		arg   string
	)

	for _, v := range val {
		if v == '(' {
			ok = true
			start = true
			continue
		}

		if ok {
			if v == ')' {
				end = true
				break
			}
			arg += string(v)

			if v == '\n' && start == true && end == false {
				fmt.Println("Error: Parenthesis not closed")
				os.Exit(1)
			}
			continue
		}

	}

	//trim the leading space
	trimmedArg := strings.TrimSpace(arg)

	//to == "" means it must be variable in the argument
	if to == "" {
		if strings.Contains(trimmedArg, ",") {

			splitted := strings.Split(trimmedArg, ",")

			if strings.TrimSpace(strings.ToLower(splitted[1])) != "first" {
				fmt.Printf("Invalid argument %v, second argument first is allowed\n", splitted[1])
				os.Exit(1)
			} else {
				if _, ok := d.Vars[splitted[0]]; !ok {
					fmt.Println("Error: Undefined variable \"", splitted[1], "\"")
					os.Exit(1)
				}

				//make first character upper-case
				variable := strings.Title(d.Vars[splitted[0]])
				d.Vars[splitted[0]] = variable
				return
			}

		}

		//if arg first is not provided make everything upper-case of the variable value
		if _, ok := d.Vars[arg]; !ok {
			fmt.Println("Error: Undefined variable \"", arg, "\"")
			os.Exit(1)
		}

		//make first character upper-case
		variable := strings.ToUpper(d.Vars[arg])
		d.Vars[arg] = variable

	}

	// if to == "" && arg[0] == '"' {

	// 	if arg[len(arg)-1] != '"' {
	// 		fmt.Println("Error: string must be closed with double quote")
	// 		os.Exit(1)
	// 	}

	// 	trimmed := arg[1 : len(arg)-1]
	// 	variable := strings.ToUpper(trimmed)
	// 	// return variable

	// }

	if to != "" {

		if arg[0] != '"' || arg[len(arg)-1] != '"' {
			fmt.Println("Error: string must be closed with double quote")
			os.Exit(1)
		}

		if _, ok := d.Vars[to]; !ok {
			fmt.Println("Error: Undefined variable \"", arg, "\"")
			os.Exit(1)
		}

		trimmed := arg[1 : len(arg)-1]
		variableValue := strings.ToUpper(trimmed)

		d.Vars[to] = variableValue
		// return variableValue
	}

}

func split(val string) []string {
	// pos := 0
	splitted := strings.Split(val, " ")
	fields := []string{}

	var (
		ok = false
		s  string
	)

	for _, v := range splitted {

		if strings.HasPrefix(v, "\"") {

			if strings.HasSuffix(v, "\"") {
				fields = append(fields, v)
				continue
			}

			ok = true

			s += v

			continue
		}

		if ok {
			s += " "
			if strings.HasSuffix(v, "\"") {
				// end = true
				ok = false

				s += v
				fields = append(fields, s)
				continue
			}
			s += v
			continue

		}

		if v == "" {
			continue
		}
		fields = append(fields, v)

	}

	return fields

}
