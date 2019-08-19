package handler

import (
	"strconv"
	"strings"
	"time"
)

func DateToStr(date, format string) (string, error) {

	var t time.Time
	var err error
	var dateInt int64

	if strings.HasPrefix(date, "\"") {
		date = date[1 : len(date)-1]
		if strings.Contains(date, "/") {

			t, err = time.Parse("01/02/2006", date)

			if err != nil {
				return "", err
			}
		} else if strings.Contains(date, "-") {

			layout := "2006-01-02 15:04:05"

			t, err = time.Parse(layout, date)

			if err != nil {
				return "", err
			}
		}

	} else {
		dateInt, err = strconv.ParseInt(date, 10, 64)

		if err != nil {
			return "", err
		}
	}

	trimmedFormat := strings.TrimSpace(format)
	trimmedFormat = trimmedFormat[1 : len(trimmedFormat)-1]

	splittedFormats := strings.Split(trimmedFormat, "-")

	formatStr := ""

	// fmtStr := ""

	// for _, v := range trimmedFormat {

	// 	switch v {
	// 	case "m":

	// 	}

	// }

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

	var formattedDate string

	if dateInt > 0 {

	} else {
		formattedDate = t.Format(formatStr)
	}

	return formattedDate, nil

}
