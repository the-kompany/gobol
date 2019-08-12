package handler

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

//Shift parse the argument from UPSHIFT()
//then it makes the provided string uppercase or only make first character of the word uppercase
//if first is found in argument with separated by comma
func (d *Data) Shift(val, first string, shiftType int) (string, error) {

	//TODO this function should only return the uppercase string and

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

	arg0Trimmed := strings.TrimSpace(argSplitted[0])

	if len(argSplitted) > 1 {
		if strings.TrimSpace(strings.ToLower(argSplitted[1])) != "first" {
			err := fmt.Errorf("Invalid argument %v, second argument first is allowed", argSplitted[1])
			return "", err
		}

		if strings.Contains(arg0Trimmed, "\"") {

			if !strings.HasPrefix(arg0Trimmed, "\"") || !strings.HasSuffix(arg0Trimmed, "\"") {
				err := errors.New("Error: string must be closed with double quote")
				return "", err
			}

			variableValue := strings.Title(arg0Trimmed)
			// d.Vars[to] = variableValue
			return variableValue[1 : len(variableValue)-1], nil

		}

		if _, ok := d.Vars[arg0Trimmed]; !ok {
			err := fmt.Errorf("Error: Undefined variable \"%v\"", arg0Trimmed)
			return "", err
		}

		variableValue := strings.Title(d.Vars[arg0Trimmed])
		// d.Vars[to] = variableValue
		return variableValue, nil

	}

	if strings.Contains(arg0Trimmed, "\"") {
		if !strings.HasPrefix(arg0Trimmed, "\"") || !strings.HasSuffix(arg0Trimmed, "\"") {
			err := errors.New("Error: string must be closed with double quote")
			return "", err
		}

		if shiftType == 1 {
			variableValue := strings.ToUpper(arg0Trimmed)
			return variableValue[1 : len(variableValue)-1], nil

		} else if shiftType == 0 {

			variableValue := strings.ToLower(arg0Trimmed)
			return variableValue[1 : len(variableValue)-1], nil
		}

	}

	//it's a variable will assign to another variable (value of TO)
	if _, ok := d.Vars[arg0Trimmed]; !ok {
		err := fmt.Errorf("Error: Undefined variable \"%v\"", arg0Trimmed)
		return "", err
	}

	if shiftType == 1 {
		variableValue := strings.ToUpper(d.Vars[arg0Trimmed])
		// d.Vars[to] = variableValue
		return variableValue, nil

	}
	variableValue := strings.ToLower(d.Vars[arg0Trimmed])
	return variableValue, nil

}
