package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	KeyWords = []string{}
)

type parser struct {
}

type Tree struct {
}

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

func ValidPerformBlock(val string) ([]string, error) {

	trimmed := strings.TrimSpace(val)

	splitted := strings.Split(trimmed, " ")

	if splitted[2] == strings.ToLower("times") {
		_, err := strconv.Atoi(strings.TrimSpace(splitted[1]))

		if err != nil {
			return splitted, errors.New("Error: loop times must be number")
		}

	}

	if strings.ToLower(strings.TrimSpace(splitted[len(splitted)-1])) != "end-perform" {
		return splitted, errors.New("Error: Perform block must be end with END-PERFORM")
	}

	sl := []string{}

	for _, v := range splitted {
		if len(strings.TrimSpace(v)) < 1 {
			continue
		} else {
			str := strings.TrimSpace(v)
			sl = append(sl, str)
		}
	}

	for _, v := range sl {
		if strings.HasSuffix(v, "\n") {
			fmt.Println("new")
		}
		// fmt.Println(len(v))
	}

	return sl, nil

}
