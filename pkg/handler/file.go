package handler

import "fmt"

func (d *Data) Close(fileHandle string) {

	if v, ok := d.FileData[fileHandle]; !ok {

		fmt.Println("Undefined file handle, ", fileHandle)
	} else {
		v.Close()
	}

}
