package handler

import "testing"

func TestMove(t *testing.T) {

	d := &Data{}
	d.Vars = make(map[string]interface{})
	d.Vars["VAR2"] = "test data"
	d.Vars["VAR1"] = "value1"
	str := "MOVE UPSHIFT(VAR2) TO VAR1"

	d.Move(str)

	if d.Vars["VAR1"] != "TEST DATA" {
		t.Errorf("Expected %v got %v", "TEST DATA", d.Vars["VAR1"])
	}

	//test move with date2str function

	// dateStr := "MOVE DATE2STR(\"08/07/2019\", \"yy-mm-dd\") TO VAR1"

	// d.Move(dateStr)

	// expected := "2019-Aug-07"
	// if d.Vars["VAR1"] != expected {
	// 	t.Errorf("Expected %v, got %v", expected, d.Vars["VAR1"])
	// }

}
