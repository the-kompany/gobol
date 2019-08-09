package handler

import "testing"

func TestIfBlock(t *testing.T) {

	str := "IF \"gobol\" = \"gobol\" THEN   DISPLAY \"This is true\" END-IF"

	d := &Data{}
	d.IfBlock(str)

}
