package handler

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/the-kompany/gobol/utils"
)

func (d *Data) Move(val string) {

	trimmed := strings.TrimSpace(val)

	splitted := utils.Split(trimmed)

	toLow := strings.ToLower(splitted[2])
	if toLow != "to" || len(splitted) < 4 {
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
