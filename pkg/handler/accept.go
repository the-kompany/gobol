package handler

import "fmt"

func (d *Data) Accept(val string) string {

	var in string

	fmt.Scanf("%v", &in)

	return in

}
