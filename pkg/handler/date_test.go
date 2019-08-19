package handler

import (
	"testing"
)

func TestDate2Str(t *testing.T) {

	result, err := DateToStr("\"08/07/2019\"", "\"yy-mm-dd\"")

	if err != nil {
		t.Errorf("%v", err)
	}

	expected := "2019-Aug-07"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	result, err = DateToStr("\"08/07/2019\"", "\"dd-mm-yy\"")

	if err != nil {
		t.Errorf("%v", err)
	}

	expected = "07-Aug-2019"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	//test mysql date-time

	result, err = DateToStr("\"2019-01-15 20:05:25\"", "\"yy-mm-dd\"")

	if err != nil {
		t.Errorf("%v", err)
	}

	expected = "2019-Jan-15"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

}
