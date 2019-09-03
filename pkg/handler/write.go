package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func (d *Data) csvWrite(tokens []string, csvWriter *csv.Writer, writeCount int) {

	writer := csvWriter

	//find the record
	recordData := d.Record[tokens[1]]

	rows := []string{}
	headers := []string{}

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

	err := writer.Write(rows)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func (d *Data) csvToJasonWrite(tokens []string, writeCount int, file *os.File) {

	//find the record
	recordData := d.Record[tokens[1]]

	var jsonRecords = make([]map[string]string, 0)

	for {
		decoder := json.NewDecoder(file)

		err := decoder.Decode(&jsonRecords)

		if err != nil {
			if err.Error() == "EOF" {

				break
			}

			fmt.Println("Error decoding JSON file", err)
			os.Exit(1)
		}

	}

	jsonRecords = append(jsonRecords, recordData)

	// create new indented JSON with the added user
	newBytes, err := json.MarshalIndent(jsonRecords, "", "    ")
	if err != nil {
		panic(err)
	}

	// write the new indented json over original file
	_, err = file.WriteAt(newBytes, 0)
	if err != nil {
		panic(err)
	}

	// content, err := ioutil.ReadFile("output.json")

	// m := make([]map[string]string, 0)

	// if err = json.Unmarshal(content, &m); err != nil {

	// 	log.Fatal(err)

	// }

}

func (d *Data) csvToFixedWrite(tokens []string, writeCount int) {

}
