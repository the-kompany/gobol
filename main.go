package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	//TODO get the filename from argument

	f, err := os.Open("hello.gbl")

	if err != nil {
		log.Fatal(err)
	}
	var extension = filepath.Ext("hello.gbl")

	if extension != ".gbl" {
		fmt.Printf("%s", "not a valid gobol file")
	}

	//temporary for testing > use scanner instead?
	b := make([]byte, 1024)
	f.Read(b)

	tkn := strings.Fields(string(b))

	fmt.Println(tkn)

}
