package main

import (
	"testing"
)

func TestUpShift(t *testing.T) {

	d := &Data{}
	//UPSHIFT(VAR)
	d.Vars = make(map[string]string)
	d.Vars["var1"] = "gobol"
	// v := "UPSHIFT( \"var1\")"
	v := "UPSHIFT(var1)"
	d.handleUpShift(v, "", "")

	if d.Vars["var1"] != "GOBOL" {
		t.Errorf("Expected %v, got %v", "GOBOL", d.Vars["var1"])
	}

	//test for first character
	d.Vars["var1"] = "gobol"
	v = "UPSHIFT(var1, first)"
	d.handleUpShift(v, "", "")

	if d.Vars["var1"] != "Gobol" {
		t.Errorf("Expected %v, got %v", "Gobol", d.Vars["var1"])
	}

}

func TestMoveWithShift(t *testing.T) {

}
