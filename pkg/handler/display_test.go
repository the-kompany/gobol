package handler

import "testing"

// func TestDisplay(t *testing.T) {

// 	d := Data{}
// 	d.Vars = make(map[string]interface{})

// 	str := "DISPLAY \"ok something\""
// 	d.Display(str)
// }

func BenchmarkDisplay(b *testing.B) {

	d := Data{}
	d.Vars = make(map[string]interface{})

	str := "DISPLAY UPSHIFT(\"ok something\")"
	for n := 0; n < b.N; n++ {

		d.Display(str)
	}

}
