package handler

import (
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

	if strings.HasPrefix(strings.ToLower(splitted[1]), "upshift") {
		d.UpShift(splitted[1], "", splitted[3])
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
