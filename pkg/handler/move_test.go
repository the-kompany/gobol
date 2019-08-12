package handler

import "testing"

func TestMove(t *testing.T) {

	d := &Data{}
	d.Vars = make(map[string]string)
	d.Vars["VAR2"] = "test data"
	d.Vars["VAR1"] = "value1"
	str := "MOVE UPSHIFT(VAR2) TO VAR1"

	d.Move(str)

	if d.Vars["VAR1"] != "TEST DATA" {
		t.Errorf("Expected %v got %v", "TEST DATA", d.Vars["VAR1"])
	}
}
