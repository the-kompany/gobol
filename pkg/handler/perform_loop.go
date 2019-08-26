package handler

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (d *Data) PerformLoopBlock(tokens []string) {

	if strings.ToLower(tokens[2]) == "times" {
		times, _ := strconv.Atoi(tokens[1])

		for i := 0; i < times; i++ {
			for i := 3; i < len(tokens)-1; i++ {

				trimmed := strings.TrimSpace(tokens[i])

				switch strings.ToLower(trimmed) {

				case "display":
					var actionStr string

					if strings.HasPrefix(strings.TrimSpace(tokens[i+1]), "\"") {

						actionStr = trimmed + " " + tokens[i+1]

					} else {
						actionStr = trimmed + " " + tokens[i+1]

					}

					d.Display(actionStr)

				case "move":
					var actionStr string
					pos := i
					actionStr += trimmed
					for strings.ToLower(tokens[pos]) != "to" {
						pos++
						actionStr += " " + tokens[pos]

					}

					actionStr += " " + tokens[pos+1]
					d.Move(actionStr)

				}
			}

		}
	}

	// log.Println("tokens ", tokens[4])

	if strings.ToLower(tokens[1]) == "read" {

		var fileReference = tokens[2]
		var recordName = tokens[4]

		// log.Println("file ref ", d.File[fileReference][1][0])
		// log.Println("record name", recordName)
		// log.Println("data ", d.Record[recordName]["Title545"])

		for i := 0; i < len(d.File[fileReference]); i++ {
			// log.Println("debug", d.File[fileReference])

			// fr := d.File[fileReference][0][0]

			//first line is title in csv
			//start the iteration from 1
			if i > 0 {

				for k, v := range d.File[fileReference][i] {

					frTitle := d.File[fileReference][0][k]

					if _, ok := d.Record[recordName][frTitle]; !ok {

						fmt.Println("file reference does not match with the record definition", frTitle)
						os.Exit(1)
					}

					d.Record[recordName][frTitle] = v
				}
			}

			for i := 3; i < len(tokens)-1; i++ {

				trimmed := strings.TrimSpace(tokens[i])

				switch strings.ToLower(trimmed) {

				case "display":
					var actionStr string

					if strings.HasPrefix(strings.TrimSpace(tokens[i+1]), "\"") {

						actionStr = trimmed + " " + tokens[i+1]

					} else {
						actionStr = trimmed + " " + tokens[i+1]

					}

					d.Display(actionStr)

				case "move":
					var actionStr string
					pos := i
					actionStr += trimmed
					for strings.ToLower(tokens[pos]) != "to" {
						pos++
						actionStr += " " + tokens[pos]

					}

					actionStr += " " + tokens[pos+1]
					d.Move(actionStr)

				}
			}

		}
	}

	if strings.ToLower(tokens[1]) == "until" {

		if tokens[3] == "=" {
			var leftValue string
			var rightValue string
			var leftVar bool
			var rightVar bool

			trimmedVal := strings.TrimSpace(tokens[2])

			if trimmedVal[0] != '"' && !isNumeric(trimmedVal) {
				if _, ok := d.Vars[trimmedVal]; !ok {
					fmt.Println("Error: Undefined variable \"", trimmedVal, "\"")
					os.Exit(1)
				}

				leftValue = d.Vars[trimmedVal]
				leftVar = true
			} else {
				leftValue = trimmedVal
			}

			trimmedVal = strings.TrimSpace(tokens[4])

			if trimmedVal[0] != '"' && !isNumeric(trimmedVal) {
				if _, ok := d.Vars[trimmedVal]; !ok {
					fmt.Println("Error: Undefined variable \"", trimmedVal, "\"")
					os.Exit(1)
				}

				rightValue = d.Vars[trimmedVal]
				rightVar = true
			} else {
				rightValue = trimmedVal
			}

			for leftValue != rightValue {
				for i := 5; i < len(tokens)-1; i++ {

					d.executeActionBlock(tokens, i)
				}

				if leftVar {

					leftValue = d.Vars[strings.TrimSpace(tokens[2])]
				}

				if rightVar {
					leftValue = d.Vars[strings.TrimSpace(tokens[4])]
				}
			}

		}
	}

	if strings.ToLower(tokens[1]) == "varying" {
		i, err := strconv.Atoi(tokens[4])

		if err != nil {
			fmt.Println("Counter variable value must be an integer, given: ", tokens[4])
		}

		//create the counter variable

		counterVar := make(map[string]int)

		counterVar[tokens[2]] = i

		var rightVar int

		if !isNumeric(tokens[10]) {
			if v, ok := d.Vars[tokens[10]]; ok {
				rightVar, _ = strconv.Atoi(v)
			} else {
				fmt.Println("Error undefined ", tokens[10])
				os.Exit(1)

			}
		} else {
			rightVar, _ = strconv.Atoi(tokens[10])
		}

		incrementValue, err := strconv.Atoi(tokens[4])

		if err != nil {

		}

		// log.Println(counterVar[tokens[2]], rightVar)

		d.Vars[tokens[2]] = tokens[4]

		for !untilValid(tokens[9], counterVar[tokens[2]], rightVar) {

			d.Vars[tokens[2]] = strconv.Itoa(counterVar[tokens[2]])
			counterVar[tokens[2]] += incrementValue

			for i := 11; i < len(tokens)-1; i++ {

				d.executeActionBlock(tokens, i)
			}
		}

	}

}

func untilValid(operator string, value1, value2 int) bool {

	switch operator {
	case "=":
		if value1 == value2 {
			return true
		}
		return false

	case ">":
		if value1 > value2 {
			return true
		}

		return false

	case ">=":
		if value1 >= value2 {
			return true
		}
		return false

	case "<":

		if value1 < value2 {
			return true
		}
		return false

	}

	return false
}

func (d *Data) executeActionBlock(tokens []string, i int) {

	trimmed := strings.TrimSpace(tokens[i])

	switch strings.ToLower(trimmed) {

	case "display":
		var actionStr string

		if strings.HasPrefix(strings.TrimSpace(tokens[i+1]), "\"") {

			actionStr = trimmed + " " + tokens[i+1]

		} else {
			actionStr = trimmed + " " + tokens[i+1]

		}

		d.Display(actionStr)

	case "move":
		var actionStr string
		pos := i
		actionStr += trimmed
		for strings.ToLower(tokens[pos]) != "to" {
			pos++
			actionStr += " " + tokens[pos]

		}

		actionStr += " " + tokens[pos+1]
		d.Move(actionStr)
	}

}
