package handler

import "os"

type Data struct {
	Vars     map[string]interface{}
	Line     int
	Lines    []Token
	Record   map[string]map[string]string
	File     map[string][][]string
	FileData map[string]*os.File
}

type Token struct {
	Value string
	Line  int
}

// type File struct {
// 	Data map[string][][]string
// }
