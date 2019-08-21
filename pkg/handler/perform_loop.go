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
					actionStr := trimmed + " " + tokens[i+1] + " " + tokens[i+2] + " " + tokens[i+3]
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
					actionStr := trimmed + " " + tokens[i+1] + " " + tokens[i+2] + " " + tokens[i+3]
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

				if leftVar {

					leftValue = d.Vars[strings.TrimSpace(tokens[2])]
				}

				if rightVar {
					leftValue = d.Vars[strings.TrimSpace(tokens[4])]
				}
			}

		}
	}

}
