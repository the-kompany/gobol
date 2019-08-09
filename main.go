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
	Vars  map[string]string
	Line  int
	Lines []string
}

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
	d := &Data{}
	d.Vars = make(map[string]string)

	//scan eeach line and append to slice
	//better for parsing
	for scanner.Scan() {

		d.Lines = append(d.Lines, scanner.Text())

	}

	//TODO Parse it

	for _, v := range d.Lines {
		if strings.HasPrefix(v, "move") || strings.HasPrefix(v, "MOVE") {
			d.handleMove(v)
			continue
		}

		if strings.HasPrefix(v, "display") || strings.HasPrefix(v, "DISPLAY") {
			d.handleDisplay(v)
		}

		if strings.HasPrefix(v, "upshift") || strings.HasPrefix(v, "UPSHIFT") {
			d.handleUpShift(v, "", "")
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

	if strings.HasPrefix(strings.ToLower(splitted[1]), "upshift") {
		d.handleUpShift(splitted[1], "", splitted[3])
	}

	if strings.HasPrefix(splitted[1], "\"") {
		d.Vars[splitted[3]] = splitted[1]
	} else {
		if _, ok := d.Vars[splitted[3]]; !ok {
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

//handleUpShift parse the argument from UPSHIFT()
//then it makes the provided string uppercase or only make first character of the word uppercase
//if first is found in argument with separated by comma
func (d *Data) handleUpShift(val, first, to string) string {

	//TODO this function should only return the uppercase string and
	//it should not assin the value to the variable
	//Refactor needed for this: remove the to and first argument

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
	argSplitted := strings.Split(trimmedArg, ",")

	//to == "" means it must be variable in the argument
	if to == "" {
		if len(argSplitted) > 1 {

			if strings.TrimSpace(strings.ToLower(argSplitted[1])) != "first" {
				fmt.Printf("Invalid argument %v, second argument first is allowed\n", argSplitted[1])
				os.Exit(1)
			} else {
				if _, ok := d.Vars[argSplitted[0]]; !ok {
					fmt.Println("Error: Undefined variable \"", argSplitted[1], "\"")
					os.Exit(1)
				}

				//make first character upper-case
				variable := strings.Title(d.Vars[argSplitted[0]])
				d.Vars[argSplitted[0]] = variable
				return variable
			}

		}

		//if arg first is not provided make everything upper-case of the variable value
		if _, ok := d.Vars[argSplitted[0]]; !ok {
			fmt.Println("Error: Undefined variable \"", argSplitted[0], "\"")
			os.Exit(1)
		}

		//make first character upper-case
		variable := strings.ToUpper(d.Vars[argSplitted[0]])
		d.Vars[argSplitted[0]] = variable
		return variable

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

		if _, ok := d.Vars[to]; !ok {
			fmt.Println("Error: Undefined variable \"", arg, "\"")
			os.Exit(1)
		}

		if len(argSplitted) > 1 {
			if strings.TrimSpace(strings.ToLower(argSplitted[1])) != "first" {
				fmt.Printf("Invalid argument %v, second argument first is allowed\n", argSplitted[1])
				os.Exit(1)
			}

			if strings.HasPrefix(argSplitted[0], "\"") {
				if !strings.HasSuffix(argSplitted[0], "\"") {
					fmt.Println("Error: string must be closed with double quote")
					os.Exit(1)
				}
				variableValue := strings.Title(argSplitted[0])
				d.Vars[to] = variableValue
				return variableValue

			} else {
				//it's a variable will assign to another variable (value of TO)

				if _, ok := d.Vars[argSplitted[0]]; !ok {
					fmt.Println("Error: Undefined variable \"", arg, "\"")
					os.Exit(1)
				}

				variableValue := strings.Title(d.Vars[argSplitted[0]])
				d.Vars[to] = variableValue
				return variableValue

			}

		}
		// trimmed := arg[1 : len(arg)-1]

		//it's a string will assign to variable
		if strings.HasPrefix(argSplitted[0], "\"") {
			if !strings.HasSuffix(arg, "\"") {
				fmt.Println("Error: string must be closed with double quote")
				os.Exit(1)
			}

			variableValue := strings.ToUpper(argSplitted[0])
			d.Vars[to] = variableValue[1 : len(variableValue)-1]

		} else {
			//it's a variable will assign to another variable (value of TO)

			if _, ok := d.Vars[argSplitted[0]]; !ok {
				fmt.Println("Error: Undefined variable \"", to, "\"")
				os.Exit(1)
			}

			variableValue := strings.ToUpper(d.Vars[argSplitted[0]])
			d.Vars[to] = variableValue
			return variableValue

		}

		// return variableValue
	}

	return ""

}

func split(val string) []string {
	// pos := 0
	splitted := strings.Split(val, " ")
	fields := []string{}

	var (
		ok               = false
		s                string
		startParenthesis = false
		parenthesisStr   string
	)

	for _, v := range splitted {

		if strings.Contains(v, "(") {
			startParenthesis = true
			parenthesisStr = v
			if strings.HasSuffix(v, ")") {
				fields = append(fields, v)
				startParenthesis = false
			}
			continue

		}

		if startParenthesis {
			parenthesisStr += " "
			if strings.HasSuffix(v, ")") {
				startParenthesis = false
				parenthesisStr += v
				fields = append(fields, parenthesisStr)
				continue
			}

			parenthesisStr += v

		}

		if !startParenthesis {

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

	}

	return fields

}
