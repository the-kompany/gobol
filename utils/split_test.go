package utils

import (
	"testing"
)

func TestSplit(t *testing.T) {

	str := "MOVE \"gobol\" TO VAR2 "

	expectedTokens := []string{"MOVE", "\"gobol\"", "TO", "VAR2"}
	tokens := Split(str)

	if len(tokens) != len(expectedTokens) {
		t.Errorf("Expected lenght for token fields %v got %v ", len(expectedTokens), len(tokens))
	}

	for k, v := range tokens {
		if expectedTokens[k] != v {
			t.Errorf("Expected %v got %v", expectedTokens[k], v)
		}
	}

	//test for multiple word in string
	str = "MOVE \"gobol is great\" TO VAR2 "

	expectedTokens = []string{"MOVE", "\"gobol is great\"", "TO", "VAR2"}
	tokens = Split(str)

	if len(tokens) != len(expectedTokens) {
		t.Errorf("Expected lenght for token fields %v got %v ", len(expectedTokens), len(tokens))
	}

	for k, v := range tokens {
		if expectedTokens[k] != v {
			t.Errorf("Expected %v got %v", expectedTokens[k], v)
		}
	}

}
