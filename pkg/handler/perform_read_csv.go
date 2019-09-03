package handler

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func (d *Data) performReadCSV(tokens []string, fileReference, recordName string) {

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
					if strings.ToLower(v) == "display" || strings.ToLower(v) == "end-perform" || strings.ToLower(v) == "move" || strings.ToLower(v) == "write" || strings.ToLower(v) == "if" {
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
			case "if":
				var actionStr string
				pos := i
				actionStr += trimmed
				for strings.ToLower(tokens[pos]) != "end-if" {
					pos++
					actionStr += " " + tokens[pos]

				}

				i = pos

				d.IfBlock(actionStr)

			case "write":

				//
				//check if file name has double quote, trime it

				writeTokens := tokens[i : i+6]
				fileName := writeTokens[3]

				if strings.HasPrefix(fileName, "\"") {
					fileName = fileName[1 : len(fileName)-1]
				}

				if strings.ToLower(tokens[i+5]) == "csv" {

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

					if fileSize < 1 {

						d.csvWrite(writeTokens, csvWriter, 1)
					} else {

						d.csvWrite(writeTokens, csvWriter, 2)
					}
					csvWriter.Flush()

					//handle and write to json file
				} else if strings.ToLower(tokens[i+5]) == "json" {

					outFile, err = os.OpenFile("output.json", os.O_RDWR|os.O_CREATE, 0644)

					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					d.csvToJasonWrite(writeTokens, writeCount, outFile)

					outFile.Close()

				} else if strings.ToLower(tokens[i+5]) == "fixed" {

					if outFile == nil {

						outFile, err = os.OpenFile("output.json", os.O_RDWR|os.O_CREATE, 0644)

						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
					}

					d.csvToFixedWrite(writeTokens, writeCount, outFile)

				}
				writeCount++

			}
		}

	}
}
