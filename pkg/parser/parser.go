package parser

import (
	"strings"
)

var (
	KeyWords = []string{}
)

func ValidIfBlock(val string) bool {

	if !strings.HasSuffix(strings.ToLower(val), "end-if") {
		return false
	}

	if !strings.Contains(strings.ToLower(val), "then") {
		return false
	}

	//steps for if block
	//boolean
	// equal to
	// greater than
	// less than

	//check equal to

	return true
}

//check if the string is a valid function call

func ValidateFunctionCall(val string) {

}
