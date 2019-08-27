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

	var str string

	for _, v := range splitted[1:] {
		if strings.HasPrefix(strings.ToLower(v), "upshift") {
			upShifted, err := d.Shift(splitted[1], "", 1)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			str += upShifted

		} else if strings.HasPrefix(strings.ToLower(v), "current_datetime") {
			currentDateTime := d.CurrentDateTime()
			fmt.Println(currentDateTime)
			return
		} else if strings.HasPrefix(strings.ToLower(v), "downshift") {
			downShifted, err := d.Shift(v, "", 0)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			str += downShifted

		} else if strings.HasPrefix(strings.ToLower(v), "current_date") {
			currentDate := d.CurrentDate()

			str += currentDate
		} else if strings.HasPrefix(strings.ToLower(v), "current_time") {
			currentTime := d.CurrentTime()
			str += currentTime
			return
		} else if strings.HasPrefix(v, "\"") {

			trimmedQuote := v[1 : len(v)-1]
			str += trimmedQuote

		} else if strings.Contains(v, ".") {

			splittedRecordName := strings.Split(v, ".")

			str += d.Record[splittedRecordName[0]][splittedRecordName[1]]

		} else if _, ok := d.Vars[v]; !ok {
			fmt.Println("Error: Undefined variable \"", splitted[1], "\"")
			os.Exit(1)
		} else {
			varData := d.Vars[v]

			formatted := fmt.Sprintf("%v", varData)
			str += formatted

		}
	}

	fmt.Printf("%v\n", str)

}
