package handler

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func (d *Data) Write(tokens []string, csvWriter *csv.Writer, writeCount int) {

	fmt.Println(tokens)

	writer := csvWriter

	//find the record
	recordData := d.Record[tokens[1]]

	rows := []string{}
	headers := []string{}
	//k = header, v = value
	log.Println(recordData)
	for k, v := range recordData {

		//write the csv header
		if writeCount == 1 {
			headers = append(headers, k)

		}
		rows = append(rows, v)
	}

	//if first time then create the header
	if writeCount == 1 {

		err := writer.Write(headers)

		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	fmt.Println("rows ", rows)

	err := writer.Write(rows)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
