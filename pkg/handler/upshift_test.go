package handler

import (
	"testing"
)

func TestShift(t *testing.T) {

	d := &Data{}
	d.Vars = make(map[string]interface{})

	v := "UPSHIFT(\"gobol\")"
	upShifted, err := d.Shift(v, "", 1)

	if err != nil {
		t.Errorf("Expected %v, got %v", "GOBOL", err)
	} else if upShifted != "GOBOL" {
		t.Errorf("Expected %v, got %v", "GOBOL", upShifted)
	}

	v = "UPSHIFT(\"gobol\" , first)"
	upShifted, err = d.Shift(v, "", 1)

	if err != nil {
		t.Errorf("Expected %v, got %v", "Gobol", err)
	} else if upShifted != "Gobol" {
		t.Errorf("Expected %v, got %v", "Gobol", upShifted)
	}

	d.Vars["var1"] = "gobol"
	v = "UPSHIFT(var1)"
	upShifted, err = d.Shift(v, "", 1)

	if err != nil {
		t.Errorf("Expected %v, got %v", "GOBOL", err)
	} else if upShifted != "GOBOL" {
		t.Errorf("Expected %v, got %v", "GOBOL", upShifted)
	}

	d.Vars["var1"] = "gobol"
	v = "UPSHIFT(var1, first)"
	upShifted, err = d.Shift(v, "", 1)

	if err != nil {
		t.Errorf("Expected %v, got %v", "Gobol", upShifted)
	} else if upShifted != "Gobol" {
		t.Errorf("Expected %v, got %v", "Gobol", upShifted)
	}

	//test for error
	v = "UPSHIFT(\"gobol)"
	_, err = d.Shift(v, "", 1)

	expected := "Error: string must be closed with double quote"
	if err == nil {
		t.Errorf("Expected %v, got %v", expected, err)
	}

	v = "UPSHIFT(gobol\")"
	_, err = d.Shift(v, "", 1)

	expected = "Error: string must be closed with double quote"
	if err == nil {
		t.Errorf("Expected %v, got %v", expected, err)
	}

	v = "UPSHIFT(var5)"
	_, err = d.Shift(v, "", 1)

	expected = "Error: Undefined variable \"var5\""

	if err.Error() != expected {
		t.Errorf("Expected %v, got %v", expected, err)
	}

	//------------------------ Test DownShift-----------------------

	v = "SHIFT(\"GOBOL\")"
	downShifted, err := d.Shift(v, "", 0)

	if err != nil {
		t.Errorf("Expected %v, got %v", "gobol", err)
	} else if downShifted != "gobol" {
		t.Errorf("Expected %v, got %v", "gobol", downShifted)
	}

	d.Vars["var1"] = "GOBOL"
	v = "UPSHIFT(var1)"
	downShifted, err = d.Shift(v, "", 0)

	if err != nil {
		t.Errorf("Expected %v, got %v", "GOBOL", err)
	} else if downShifted != "gobol" {
		t.Errorf("Expected %v, got %v", "gobol", downShifted)
	}

}

func TestMoveWithShift(t *testing.T) {
	//
}
