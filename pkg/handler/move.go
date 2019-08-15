package handler

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/the-kompany/gobol/utils"
)

func (d *Data) Move(val string) {

	trimmed := strings.TrimSpace(val)

	splitted := utils.Split(trimmed)

	toLow := strings.ToLower(splitted[len(splitted)-2])

	if toLow != "to" {
		fmt.Println("Error: Inavalid syntax for MOVE")
		os.Exit(1)
	}

	funcName := splitted[1]
	funcNameLower := strings.ToLower(funcName)

	switch {
	case strings.HasPrefix(funcNameLower, "upshift"):
		upShifted, err := d.Shift(funcName, "", 1)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if _, ok := d.Vars[splitted[3]]; !ok {
			fmt.Println("Error: Undefined variable \"", funcName, "\"")
			os.Exit(1)
		}

		d.Vars[splitted[3]] = upShifted

	case strings.HasPrefix(funcNameLower, "downshift"):
		downShifted, err := d.Shift(funcName, "", 0)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if _, ok := d.Vars[splitted[3]]; !ok {
			fmt.Println("Error: Undefined variable \"", funcName, "\"")
			os.Exit(1)
		}

		d.Vars[splitted[3]] = downShifted

	case strings.HasPrefix(funcNameLower, "accept"):
		arg, err := getFuncArg(funcName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		v := d.Accept(arg)
		if len(v) > 0 {
			if _, ok := d.Vars[splitted[3]]; !ok {
				fmt.Println("Error: Undefined variable \"", splitted[1], "\"")
				os.Exit(1)
			}
			d.Vars[splitted[3]] = v
		}
	default:

		var (
			value      string
			leftValue  string
			rightValue string
			operator   string
		)
		valueSplitted := splitted[1 : len(splitted)-2]

		if len(valueSplitted) > 1 {
			for k, v := range valueSplitted {
				trimmedVal := strings.TrimSpace(v)

				if k == 0 {
					trimmedVal := strings.TrimSpace(valueSplitted[0])

					if trimmedVal[0] != '"' && !isNumeric(trimmedVal) {
						if _, ok := d.Vars[trimmedVal]; !ok {
							fmt.Println("Error: Undefined variable \"", trimmedVal, "\"")
							os.Exit(1)
						}

						leftValue = d.Vars[trimmedVal]
					} else {
						leftValue = trimmedVal
					}
				}

				if k == 1 {
					if trimmedVal != "+" && trimmedVal != "-" && trimmedVal != "*" && trimmedVal != "/" {
						fmt.Println("Error: Invalid syntax")
						os.Exit(1)
					}

					operator = trimmedVal

				}

				if k == 2 {
					trimmedVal := strings.TrimSpace(trimmedVal)

					if trimmedVal[0] != '"' && !isNumeric(trimmedVal) {
						if _, ok := d.Vars[trimmedVal]; !ok {
							fmt.Println("Error: Undefined variable \"", trimmedVal, "\"")
							os.Exit(1)
						}

						rightValue = d.Vars[trimmedVal]
					} else {
						rightValue = trimmedVal
					}

				}

			}

			if !isNumeric(leftValue) && !isNumeric(rightValue) {
				fmt.Println("Error: invalid syntax for move")
				os.Exit(1)
			}

			leftValueInt, _ := strconv.Atoi(leftValue)
			rightValueInt, _ := strconv.Atoi(rightValue)

			var valueInt int
			switch operator {
			case "+":

				valueInt = leftValueInt + rightValueInt

				log.Println(leftValue, rightValue, valueInt)
			case "-":
				valueInt = leftValueInt - rightValueInt

			case "*":
				valueInt = leftValueInt * rightValueInt

			case "/":
				valueInt = leftValueInt / rightValueInt

			}

			value = strconv.Itoa(valueInt)

		} else {

			trimmedVal := strings.TrimSpace(valueSplitted[0])

			if trimmedVal[0] != '"' && !isNumeric(trimmedVal) {
				if _, ok := d.Vars[trimmedVal]; !ok {
					fmt.Println("Error: Undefined variable \"", trimmedVal, "\"")
					os.Exit(1)
				}

				value = d.Vars[trimmedVal]
			} else {
				value = trimmedVal
			}

		}

		d.Vars[splitted[len(splitted)-1]] = value
	}

}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func getFuncArg(val string) (string, error) {

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
				err := errors.New("Error: Parenthesis not closed")
				return arg, err
			}
			continue
		}

	}

	return arg, nil
}
