package handler

import (
	"fmt"
	"os"
	"strings"
)

//IfBlock execute if block
func (d *Data) IfBlock(val string) {

	trimmed := strings.TrimSpace(val)
	thenSplit := strings.Split(trimmed, "THEN")

	//remove the if from string
	comparisonStr := thenSplit[0][2:]
	comparisonStr = strings.TrimSpace(comparisonStr)

	compariosnSl := []string{}
	//split it by equal sign if has = sign
	if strings.Contains(comparisonStr, "=") {
		compariosnSl = strings.Split(comparisonStr, "=")
	}

	//if comparison is true then execute the 'then' action

	var leftValue string
	var rightValue string

	if !strings.HasPrefix(strings.TrimSpace(compariosnSl[0]), "\"") {
		if _, ok := d.Vars[strings.TrimSpace(compariosnSl[0])]; !ok {
			fmt.Println("Error: undefined ", compariosnSl[0])
			os.Exit(1)
		}

		leftValue = d.Vars[strings.TrimSpace(compariosnSl[0])].(string)

	} else {
		leftValue = strings.TrimSpace(compariosnSl[0])
	}

	if !strings.HasPrefix(strings.TrimSpace(compariosnSl[1]), "\"") {
		if _, ok := d.Vars[strings.TrimSpace(compariosnSl[1])]; !ok {
			fmt.Println("Error: undefined ", compariosnSl[1])
			os.Exit(1)
		}

		rightValue = d.Vars[strings.TrimSpace(compariosnSl[1])].(string)

	} else {
		rightValue = strings.TrimSpace(compariosnSl[1])
	}

	if leftValue == rightValue {
		// d.Display("DISPLAY \"OK\"")

		thenAction := strings.Split(thenSplit[1], "END")
		thenActionTrimmed := strings.TrimSpace(thenAction[0])

		//split the function name
		//TODO there must be an execute function for executing function
		//like this: execute("DISPLAY VAR1")

		if strings.Contains(thenActionTrimmed, "DISPLAY") {
			d.Display(thenActionTrimmed)
		}

		if strings.Contains(thenActionTrimmed, "MOVE") {
			d.Move(thenActionTrimmed)
		}
	}

	// fmt.Println(thenActionTrimmed)

}
