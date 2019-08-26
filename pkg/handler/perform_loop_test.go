package handler

import "testing"

var testUntilValidData = []struct {
	operator   string
	leftVal    int
	rightValue int
	expected   bool
}{
	{"=", 1, 2, true},  //return true  for contuning the loop
	{"=", 2, 2, false}, //returns false for breaking the loop
	{">", 2, 1, true},
	{">", 1, 2, false},
	{"<", 1, 2, true},
}

func TestPerformLoop(t *testing.T) {

}

func TestUntilValid(t *testing.T) {

	for _, v := range testUntilValidData {
		got := untilValid(v.operator, v.leftVal, v.rightValue)

		if v.expected != got {
			t.Errorf("Expected %v, got %v", v.expected, got)
		}
	}
}
