package handler

import (
	"fmt"
	"os"
	"strings"

	"github.com/the-kompany/gobol/utils"
)

func (d *Data) Display(val string) {
	trimmed := strings.TrimSpace(val)

	splitted := utils.Split(trimmed)
	// log.Println(trimmed)
	if strings.HasPrefix(strings.ToLower(splitted[1]), "upshift") {
		upShifted, err := d.Shift(splitted[1], "", 1)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(upShifted)
		return

	} else if strings.HasPrefix(strings.ToLower(splitted[1]), "downshift") {
		downShifted, err := d.Shift(splitted[1], "", 0)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(downShifted)
		return
	}

	if strings.HasPrefix(splitted[1], "\"") {

		trimmedQuote := splitted[1][1 : len(splitted[1])-1]
		fmt.Println(trimmedQuote)
		return

	}

	if strings.Contains(splitted[1], ".") {

		splittedRecordName := strings.Split(splitted[1], ".")
		fmt.Println(d.Record[splittedRecordName[0]][splittedRecordName[1]])

	} else if _, ok := d.Vars[splitted[1]]; !ok {
		fmt.Println("Error: Undefined variable \"", splitted[1], "\"")
		os.Exit(1)
	} else {
		varData := d.Vars[splitted[1]]
		if strings.HasPrefix(varData, "\"") && strings.HasSuffix(varData, "\"") {
			trimmedQuote := varData[1 : len(varData)-1]
			fmt.Println(trimmedQuote)
		} else {
			fmt.Println(varData)
		}
	}

}
