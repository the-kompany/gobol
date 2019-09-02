package handler

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var writeCount int
var outFile *os.File
var csvWriter *csv.Writer

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
		r := csv.NewReader(d.FileData[fileReference])

		//read the CSV header
		header, err := r.Read()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for {

			rec, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println("Error reading file", err.Error())
				os.Exit(1)
			}

			for k, v := range header {
				if _, ok := d.Record[recordName][v]; !ok {

					fmt.Println("file reference does not match with the record definition", v)
					os.Exit(1)
				}
				d.Record[recordName][v] = rec[k]
			}

			for i := 3; i < len(tokens)-1; i++ {

				trimmed := strings.TrimSpace(tokens[i])

				switch strings.ToLower(trimmed) {

				case "display":
					var actionStr string

					actionStr += trimmed

					for _, v := range tokens[i+1:] {
						//TODO this needed to be fixed, should be handled by newline \n
						if strings.ToLower(v) == "display" || strings.ToLower(v) == "end-perform" || strings.ToLower(v) == "move" || strings.ToLower(v) == "write" {
							break
						} else {
							actionStr += " " + v
						}
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
				case "write":

					//
					//check if file name has double quote, trime it

					writeTokens := tokens[i : i+6]
					fileName := writeTokens[3]

					if strings.HasPrefix(fileName, "\"") {
						fileName = fileName[1 : len(fileName)-1]
					}

					if writeCount < 1 {

						outFile, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
						csvWriter = csv.NewWriter(outFile)

					}

					info, _ := outFile.Stat()
					fileSize := info.Size()

					writeCount++

					if fileSize < 1 {

						d.Write(writeTokens, csvWriter, 1)
					} else {

						d.Write(writeTokens, csvWriter, 2)
					}
					csvWriter.Flush()

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

				leftValue = d.Vars[trimmedVal].(string)
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

				rightValue = d.Vars[trimmedVal].(string)
				rightVar = true
			} else {
				rightValue = trimmedVal
			}

			for leftValue != rightValue {
				for i := 5; i < len(tokens)-1; i++ {

					d.executeActionBlock(tokens, i)
				}

				if leftVar {

					leftValue = d.Vars[strings.TrimSpace(tokens[2])].(string)
				}

				if rightVar {
					leftValue = d.Vars[strings.TrimSpace(tokens[4])].(string)
				}
			}

		}
	}

	if strings.ToLower(tokens[1]) == "varying" {
		i, err := strconv.Atoi(tokens[4])

		if err != nil {
			fmt.Println("Counter variable value must be an integer, given: ", tokens[4])
			os.Exit(1)
		}

		//create the counter variable

		// counterVar := make(map[string]int)

		// counterVar[tokens[2]] = i

		d.Vars[tokens[2]] = i

		var rightVar int

		if !isNumeric(tokens[10]) {
			if v, ok := d.Vars[tokens[10]]; ok {
				rightVar, _ = v.(int)
			} else {
				fmt.Println("Error undefined ", tokens[10])
				os.Exit(1)

			}
		} else {
			rightVar, _ = strconv.Atoi(tokens[10])
		}

		incrementValue, err := strconv.Atoi(tokens[6])

		if err != nil {
			fmt.Println("Error: Increment value must be an integer")
			os.Exit(1)
		}

		for !untilValid(tokens[9], d.Vars[tokens[2]].(int), rightVar) {

			for i := 11; i < len(tokens)-1; i++ {

				d.executeActionBlock(tokens, i)
			}
			d.Vars[tokens[2]] = d.Vars[tokens[2]].(int) + incrementValue

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

		actionStr += trimmed

		for _, v := range tokens[i+1:] {

			if strings.ToLower(v) == "display" || strings.ToLower(v) == "end-perform" || strings.ToLower(v) == "move" {
				break
			} else {
				actionStr = actionStr + " " + v
			}
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
