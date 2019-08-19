package handler

import (
	"fmt"
	"os"
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

	trimmedFormat = strings.ToLower(trimmedFormat)

	yCount := 0
	mCount := 0
	dCount := 0
	formattedStr := ""

	isDay := false

	for k, v := range trimmedFormat {
		switch v {
		case 'y':
			if isDay {
				continue
			}
			yCount++
			if yCount == 1 {
				if trimmedFormat[k+1] != 'y' {
					yCount = 0
					fmt.Println("Error: invalid date format for year")
					os.Exit(1)
				}
				continue
			}

			if len(trimmedFormat)-1 == k || trimmedFormat[k+1] != 'y' {

				if yCount == 2 {
					formattedStr += "06"
					yCount = 0
				} else if yCount == 4 {
					formattedStr += "2006"
					yCount = 0
				} else {
					if trimmedFormat[k+1] != 'y' {
						yCount = 0
						fmt.Println("Error: invalid date format for year")
						os.Exit(1)
					}
				}
			}

		case '-':
			formattedStr += "-"
		case 'm':
			isDay = false
			mCount++
			if mCount == 1 {
				if trimmedFormat[k+1] != 'm' {
					formattedStr += "Jan"
					mCount = 0
					continue
				}
			}

			if mCount == 2 {
				formattedStr += "Jan"
				mCount = 0
			}
		case 'd':

			if len(trimmedFormat)-1 != k {

				if trimmedFormat[k+1] == 'a' {
					if trimmedFormat[k+2] == 'y' {
						formattedStr += "Mon"
						isDay = true
						continue
					}
				}
			}

			dCount++
			if dCount == 1 {
				if trimmedFormat[k+1] != 'd' {
					formattedStr += "2"
					dCount = 0
					isDay = false
					continue
				}
			}

			if dCount == 2 {
				isDay = false
				formattedStr += "02"
				dCount = 0
			}
		case ' ':
			formattedStr += " "

		}
	}

	var formattedDate string

	if dateInt > 0 {

	} else {
		formattedDate = t.Format(formattedStr)
	}

	return formattedDate, nil

}
