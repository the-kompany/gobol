package handler

type Data struct {
	Vars  map[string]string
	Line  int
	Lines []Token
}

type Token struct {
	Value string
	Line  int
}
