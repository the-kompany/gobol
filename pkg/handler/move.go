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

type argType int

const (
	str argType = iota
	variable
	fn
)

func (d *Data) Move(val string) {

	// getType(val, "MOVE")
	//move + string | variable | function + to + variable
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
		v := d.Accept(arg[0])
		if len(v) > 0 {
			if _, ok := d.Vars[splitted[3]]; !ok {
				fmt.Println("Error: Undefined variable \"", splitted[1], "\"")
				os.Exit(1)
			}
			d.Vars[splitted[3]] = v
		}
	case strings.HasPrefix(funcNameLower, "fmtdate"):
		args, err := getFuncArg(funcName)

		if err != nil {
			log.Println("Err at line ", d.Line, err)
		}

		log.Println(args)

		var inputFormat string
		if len(args) == 3 {
			inputFormat = args[1]
		}

		dateStr, err := DateToStr(d, args[0], inputFormat, args[2])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if strings.Contains(splitted[3], ".") {
			recordSplitted := strings.Split(splitted[3], ".")
			d.Record[recordSplitted[0]][recordSplitted[1]] = dateStr
		} else {
			d.Vars[splitted[3]] = dateStr
		}

	default:

		var (
			value      string
			leftValue  string
			rightValue string
			operator   string
		)
		valueSplitted := splitted[1 : len(splitted)-2]

		//check arithmetic operation with multiple value
		if len(valueSplitted) > 1 {
			for k, v := range valueSplitted {
				trimmedVal := strings.TrimSpace(v)

				if k == 0 {
					if trimmedVal[0] != '"' && !isNumeric(trimmedVal) && trimmedVal[0] != '\'' {

						fmt.Println(trimmedVal[0])
						if _, ok := d.Vars[trimmedVal]; !ok {
							fmt.Println("Error: Undefined variable \"", trimmedVal, "\"")
							os.Exit(1)
						}

						leftValue = d.Vars[trimmedVal].(string)
					} else {
						if !isNumeric(trimmedVal) {
							leftValue = trimmedVal[1 : len(trimmedVal)-1]
						} else {
							leftValue = trimmedVal
						}

					}
				}

				if k > 0 {
					if trimmedVal[0] == '"' || trimmedVal[0] == '\'' {
						if !isNumeric(trimmedVal) {
							leftValue += " " + trimmedVal[1:len(trimmedVal)-1]
						} else {
							leftValue += trimmedVal
						}
					}

					if trimmedVal == "+" || trimmedVal == "-" || trimmedVal == "*" || trimmedVal == "/" {

						operator = trimmedVal

						rightValue = strings.TrimSpace(splitted[k+2])

						log.Println(rightValue)
						if rightValue[0] != '"' && !isNumeric(rightValue) && rightValue[0] != '\'' {
							if _, ok := d.Vars[rightValue]; !ok {
								fmt.Println("Error: Undefined variable \"", trimmedVal, "\"")
								os.Exit(1)
							}

							rightValue = d.Vars[rightValue].(string)
						}

					}

				}

			}

			//if operator is found, do the mathmetical operation
			if len(operator) > 0 {

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

				case "-":
					valueInt = leftValueInt - rightValueInt

				case "*":
					valueInt = leftValueInt * rightValueInt

				case "/":
					valueInt = leftValueInt / rightValueInt

				}

				value = strconv.Itoa(valueInt)
			} else {
				value = leftValue
			}

		} else {

			trimmedVal := strings.TrimSpace(valueSplitted[0])

			if trimmedVal[0] != '"' && !isNumeric(trimmedVal) && trimmedVal[0] != '\'' {

				if strings.Contains(trimmedVal, ".") {
					recordSplitted := strings.Split(trimmedVal, ".")

					if _, ok := d.Record[recordSplitted[0]][recordSplitted[1]]; !ok {
						fmt.Println("Error: Undefined record \"", trimmedVal, "\"")
						os.Exit(1)
					}

					value = d.Record[splitted[0]][splitted[1]]
				} else {

					if _, ok := d.Vars[trimmedVal]; !ok {
						fmt.Println("Error: Undefined variable \"", trimmedVal, "\"")
						os.Exit(1)
					}

					value = d.Vars[trimmedVal].(string)
				}

			} else {
				if isNumeric(trimmedVal) {
					value = trimmedVal
				} else if strings.HasPrefix(trimmedVal, "\"") || strings.HasPrefix(trimmedVal, "'") {

					value = trimmedVal[1 : len(trimmedVal)-1]

				}
			}

		}

		varName := splitted[len(splitted)-1]
		if strings.Contains(varName, ".") {
			recordSplitted := strings.Split(varName, ".")
			d.Record[recordSplitted[0]][recordSplitted[1]] = value
		} else {
			d.Vars[splitted[len(splitted)-1]] = value
		}
	}

}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func getType(val, keyWord string) argType {

	splitted := strings.Split(val, keyWord)

	log.Println("type ", splitted[1])

	return 1
}

func getFuncArg(val string) ([]string, error) {

	var (
		ok    = false
		start = false
		end   = false
		args  []string
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
				args = append(args, arg)

				break
			}

			if v == ',' {
				args = append(args, arg)
				arg = ""
				continue
			}
			arg += string(v)

			if v == '\n' && start == true && end == false {
				err := errors.New("Error: Parenthesis not closed")
				return args, err
			}
			continue
		}

	}

	return args, nil
}
