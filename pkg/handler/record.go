package handler

import (
	"strings"
)

func (d *Data) ExecuteRecord(val string) {

	splitted := strings.Split(val, " ")

	d.Record = make(map[string]map[string]string)

	recordName := strings.TrimSpace(splitted[1])
	d.Record[recordName] = make(map[string]string)

	for _, v := range splitted[2 : len(splitted)-1] {

		trimmedFieldName := strings.TrimSpace(v)

		if len(trimmedFieldName) < 1 {
			continue
		}

		d.Record[recordName][trimmedFieldName] = ""

	}
}
