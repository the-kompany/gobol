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
