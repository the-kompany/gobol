package handler

import (
	"strconv"
	"strings"
	"time"
)

func (d *Data) DateToStr(date, format string) (string, error) {

	trimmedFormat := strings.TrimSpace(format)
	trimmedFormat = trimmedFormat[1 : len(trimmedFormat)-1]

	splittedFormats := strings.Split(trimmedFormat, "-")

	formatStr := ""

	dateInt, err := strconv.ParseInt(date, 10, 64)

	if err != nil {
		return "", err
	}

	for _, v := range splittedFormats {
		switch v {
		case "mm":
			if len(formatStr) < 1 {
				formatStr += "Jan"
			} else {
				formatStr += "-Jan"
			}
		case "yy":
			if len(formatStr) < 1 {
				formatStr += "2006"
			} else {
				formatStr += "-2006"
			}
		case "dd":
			if len(formatStr) < 1 {
				formatStr += "02"
			} else {
				formatStr += "-02"
			}
		}
	}

	formattedDate := time.Unix(dateInt, 0).Format(formatStr)

	return formattedDate, nil

}
