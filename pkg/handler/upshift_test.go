package handler

import (
	"testing"
)

func TestUpShift(t *testing.T) {

	d := &Data{}
	d.Vars = make(map[string]string)

	v := "UPSHIFT(\"gobol\")"
	upShifted, err := d.UpShift(v, "", "")

	if err != nil {
		t.Errorf("Expected %v, got %v", "GOBOL", err)
	} else if upShifted != "GOBOL" {
		t.Errorf("Expected %v, got %v", "GOBOL", upShifted)
	}

	v = "UPSHIFT(\"gobol\" , first)"
	upShifted, err = d.UpShift(v, "", "")

	if err != nil {
		t.Errorf("Expected %v, got %v", "Gobol", err)
	} else if upShifted != "Gobol" {
		t.Errorf("Expected %v, got %v", "Gobol", upShifted)
	}

	d.Vars["var1"] = "gobol"
	v = "UPSHIFT(var1)"
	upShifted, err = d.UpShift(v, "", "")

	if err != nil {
		t.Errorf("Expected %v, got %v", "GOBOL", err)
	} else if upShifted != "GOBOL" {
		t.Errorf("Expected %v, got %v", "GOBOL", upShifted)
	}

	d.Vars["var1"] = "gobol"
	v = "UPSHIFT(var1, first)"
	upShifted, err = d.UpShift(v, "", "")

	if err != nil {
		t.Errorf("Expected %v, got %v", "Gobol", upShifted)
	} else if upShifted != "Gobol" {
		t.Errorf("Expected %v, got %v", "Gobol", upShifted)
	}

	//test for error
	v = "UPSHIFT(\"gobol)"
	_, err = d.UpShift(v, "", "")

	expected := "Error: string must be closed with double quote"
	if err == nil {
		t.Errorf("Expected %v, got %v", expected, err)
	}

	v = "UPSHIFT(gobol\")"
	_, err = d.UpShift(v, "", "")

	expected = "Error: string must be closed with double quote"
	if err == nil {
		t.Errorf("Expected %v, got %v", expected, err)
	}

	v = "UPSHIFT(var5)"
	_, err = d.UpShift(v, "", "")

	expected = "Error: Undefined variable \"var5\""

	if err.Error() != expected {
		t.Errorf("Expected %v, got %v", expected, err)
	}

}

func TestMoveWithShift(t *testing.T) {
	//
}
