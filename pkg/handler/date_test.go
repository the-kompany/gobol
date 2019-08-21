package handler

import (
	"testing"
)

var dateTestCases = []struct {
	input        string
	inputFormat  string
	outputFormat string
	expected     string
}{
	{"\"19/08/15\"", "yy/mm/dd", "\"yy-mm-dd\"", "19-08-15"},
	{"\"08/07/2019\"", "mm/dd/yyyy", "\"dd-mm-yy\"", "07-08-19"},
	{"\"08-07-2019\"", "mm-dd-yyyy", "\"yy-mm-dd\"", "19-08-07"},
	{"\"08-15-2019\"", "mm-dd-yyyy", "\"dd-mm-yy\"", "15-08-19"},
	{"\"08-15-2019\"", "mm-dd-yyyy", "\"dd-mm-yyyy\"", "15-08-2019"}, // //test year 4 digit dd-mm-yyyy
	{"\"08-15-2019\"", "mm-dd-yyyy", "\"yyyy-mm-dd\"", "2019-08-15"}, //test 4 digit year in the begining
	{"\"08-15-2019\"", "mm-dd-yyyy", "\"yyyymmdd\"", "20190815"},
	{"\"08-15-2019\"", "mm-dd-yyyy", "\"YYYYMMDD\"", "20190815"},                     //test uppercase
	{"\"08-15-2019\"", "mm-dd-yyyy", "\"day dd mm\"", "Thu 15 08"},                   //day by name
	{"\"08-15-2019\"", "mm-dd-yyyy", "\"dd mm day\"", "15 08 Thu"},                   //day name at the end
	{"\"08-15-2019\"", "mm-dd-yyyy", "\"Month dd yyyy\"", "August 15 2019"},          //month full name
	{"\"08-15-2019\"", "mm-dd-yyyy", "\"yyyy dd Month\"", "2019 15 August"},          //Month in different position as output argument
	{"\"2019-01-15 20:05:25\"", "yyyy-mm-dd hh24:mi:ss", "\"hh:mi:ss\"", "08:05:25"}, //hour, minute second
	{"\"2019-01-15 20:05:25\"", "yyyy-mm-dd hh24:mi:ss", "\"dd mm yy hh:mi:ss\"", "15 01 19 08:05:25"},
	{"\"2019-01-15 20:05:25\"", "yyyy-mm-dd hh24:mi:ss", "\"Month dd yy hh:mi:ss\"", "January 15 19 08:05:25"},
	{"\"2019-01-15 20:05:25\"", "yyyy-mm-dd hh24:mi:ss", "\"Month dd yy hh24:mi:ss\"", "January 15 19 20:05:25"},
	{"\"2019-01-15 11:05:25\"", "yyyy-mm-dd hh24:mi:ss", "\"Month dd yy hh24:mi:ssPM\"", "January 15 19 11:05:25AM"},
	{"\"08/15/2019\"", "mm/dd/yyyy", "\"dd-mm-yyyy\"", "15-08-2019"},
}

var testCasesDateLayout = []struct {
	input    string
	expected string
}{
	{"mm-dd-yyyy", "01-02-2006"},
	{"mmddyyyy", "01022006"},
	{"dd-mm-yyyy", "02-01-2006"},
	//TODO add more test for edge cases
}

func TestDate2Str(t *testing.T) {

	for _, v := range dateTestCases {

		got, err := DateToStr(v.input, v.inputFormat, v.outputFormat)

		if err != nil {
			t.Errorf("Expected %v, got %v", v.expected, err)
		}

		if v.expected != got {
			t.Errorf("Expected %v, got %v", v.expected, got)
		}
	}

}

func TestGetDateLayout(t *testing.T) {

	for _, v := range testCasesDateLayout {

		got, err := getDateLayout(v.input)

		if err != nil {
			t.Errorf("Expected %v got %v", v.expected, err)
		}

		if v.expected != got {
			t.Errorf("Expected %v, got %v", v.expected, got)
		}
	}

}
