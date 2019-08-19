package handler

type Data struct {
	Vars   map[string]string
	Line   int
	Lines  []Token
	Record map[string]map[string]string
	File   map[string][][]string
}

type Token struct {
	Value string
	Line  int
}

// type File struct {
// 	Data map[string][][]string
// }
