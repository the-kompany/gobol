package transpiler

import (
	"fmt"
	"os"
	"strings"
)

//handleUpShift parse the argument from UPSHIFT()
//then it makes the provided string uppercase or only make first character of the word uppercase
//if first is found in argument with separated by comma
func (d *Data) HandleUpShift(val, first, to string) string {

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
