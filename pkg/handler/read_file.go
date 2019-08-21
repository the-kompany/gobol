package handler

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"
)

func (d *Data) Open(val string) {

	args, err := getFuncArg(val)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	arg := args[0][1 : len(args[0])-1]

	if !strings.Contains(strings.ToLower(val), "as") {
		fmt.Printf("Error: as keyword needed to reference the file data\n")
		os.Exit(1)
	}

	splitted := strings.Split(val, " ")

	referenceID := splitted[len(splitted)-1]

	if len(referenceID) < 1 {
		fmt.Printf("Error: valid reference id is needed for reading file data\n")
		os.Exit(1)
	}

	lines, err := ReadFile(arg)

	if err != nil {
		fmt.Printf("Error at line %v", err.Error())
		os.Exit(1)
	}

	d.File = make(map[string][][]string)

	d.File[referenceID] = lines

	// log.Println("file data", d.File[referenceID])
}

func ReadFile(filePath string) ([][]string, error) {

	fmt.Println(strings.TrimSpace(filePath))

	f, err := os.Open(strings.TrimSpace(filePath))

	if err != nil {
		return nil, errors.New("Error reading file, " + err.Error())
	}

	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()

	return lines, nil

}
