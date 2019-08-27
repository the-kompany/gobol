package handler

import (
	"fmt"
	"os"
	"strings"

	"github.com/the-kompany/gobol/utils"
)

func (d *Data) Display(val string) {
	trimmed := strings.TrimSpace(val)

	//display + function
	//display + variable
	//display + string
	//TODO imaplement if the argument is function call the function

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

	} else if strings.HasPrefix(strings.ToLower(splitted[1]), "current_datetime") {
		currentDateTime := d.CurrentDateTime()
		fmt.Println(currentDateTime)
		return
	} else if strings.HasPrefix(strings.ToLower(splitted[1]), "downshift") {
		downShifted, err := d.Shift(splitted[1], "", 0)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(downShifted)
		return
	} else if strings.HasPrefix(strings.ToLower(splitted[1]), "current_date") {
		currentDate := d.CurrentDate()
		fmt.Println(currentDate)
		return
	} else if strings.HasPrefix(strings.ToLower(splitted[1]), "current_time") {
		currentTime := d.CurrentTime()
		fmt.Println(currentTime)
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

		fmt.Printf("%v\n", varData)

	}

}
