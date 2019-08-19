package handler

import (
	"testing"
)

func TestDate2Str(t *testing.T) {

	result, err := DateToStr("\"08/07/2019\"", "\"yy-mm-dd\"")

	if err != nil {
		t.Errorf("%v", err)
	}

	expected := "19-Aug-07"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	result, err = DateToStr("\"08/07/2019\"", "\"dd-mm-yy\"")

	if err != nil {
		t.Errorf("%v", err)
	}

	expected = "07-Aug-19"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	//test mysql date-time

	result, err = DateToStr("\"2019-01-15 20:05:25\"", "\"yy-mm-dd\"")

	if err != nil {
		t.Errorf("%v", err)
	}

	expected = "19-Jan-15"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	//test dd-mm-yy
	result, err = DateToStr("\"2019-01-15 20:05:25\"", "\"dd-mm-yy\"")

	if err != nil {
		t.Errorf("%v", err)
	}

	expected = "15-Jan-19"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	//test year 4 digit dd-mm-yyyy
	result, err = DateToStr("\"2019-01-15 20:05:25\"", "\"dd-mm-yyyy\"")

	if err != nil {
		t.Errorf("%v", err)
	}

	expected = "15-Jan-2019"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	//test 4 digit year in the begining
	result, err = DateToStr("\"2019-01-15 20:05:25\"", "\"yyyy-mm-dd\"")

	if err != nil {
		t.Errorf("%v", err)
	}

	expected = "2019-Jan-15"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	//test for no dash in the dat eformat

	result, err = DateToStr("\"2019-01-15 20:05:25\"", "\"yyyymmdd\"")

	if err != nil {
		t.Errorf("%v", err)
	}

	expected = "2019Jan15"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	//test uppercase

	result, err = DateToStr("\"2019-01-15 20:05:25\"", "\"YYYYMMDD\"")

	if err != nil {
		t.Errorf("%v", err)
	}

	expected = "2019Jan15"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	//test day name
	result, err = DateToStr("\"2019-01-15 20:05:25\"", "\"day dd mm\"")

	if err != nil {
		t.Errorf("%v", err)
	}

	expected = "Tue 15 Jan"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	//test day name at the end
	result, err = DateToStr("\"2019-01-15 20:05:25\"", "\"dd mm day\"")

	if err != nil {
		t.Errorf("%v", err)
	}

	expected = "15 Jan Tue"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

}
