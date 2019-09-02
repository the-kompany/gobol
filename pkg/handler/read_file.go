package handler

import (
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

	f, err := os.Open(strings.TrimSpace(arg))

	if err != nil {
		err = errors.New("Error reading file, " + err.Error())

		fmt.Println(err)
		os.Exit(1)
	}

	d.FileData = make(map[string]*os.File)
	d.FileData[referenceID] = f

}

func (d *Data) ReadCSV() {

}
