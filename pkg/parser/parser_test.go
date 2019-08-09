package parser

import "testing"

func TestValidIfBlock(t *testing.T) {

	//test without 'end-if'
	ifBlock := "IF (DOWNSHIFT(mystring) = “gobol) THEN   MOVE “20190401” TO MYDATE  "

	if ValidIfBlock(ifBlock) != false {
		t.Errorf("Expected %v got %v", true, false)
	}

	//test without 'then'
	ifBlock = "IF (DOWNSHIFT(mystring) = “gobol)    MOVE “20190401” TO MYDATE  END-IF"

	if ValidIfBlock(ifBlock) != false {
		t.Errorf("Expected %v got %v", true, false)
	}

}

func TestValidateFunctionCall(t *testing.T) {

	str := "UPSHIFT()"

	ValidateFunctionCall(str)
}
