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
