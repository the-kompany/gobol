package handler

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func DateToStr(date, inputFormat, format string) (string, error) {

	var t time.Time
	var err error
	var dateInt int64

	if len(inputFormat) > 0 {
		log.Println(inputFormat)

		date = date[1 : len(date)-1]
		if inputFormat[0] == 'd' && inputFormat[1] == 'd' {
			t, err = time.Parse("02/01/2006", date)

			if err != nil {
				return "", err
			}
		} else if inputFormat[0] == 'm' && inputFormat[1] == 'm' {
			t, err = time.Parse("01/02/2006", date)

			if err != nil {
				return "", err
			}

		}

	} else {
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
	}

	trimmedFormat := strings.TrimSpace(format)
	trimmedFormat = trimmedFormat[1 : len(trimmedFormat)-1]

	trimmedFormat = strings.ToLower(trimmedFormat)

	yCount := 0
	mCount := 0
	dCount := 0
	hCount := 0
	sCount := 0
	formattedStr := ""

	isDay := false
	// isMonth := false

	for i := 0; i < len(trimmedFormat); i++ {
		switch trimmedFormat[i] {
		case 'y':
			if isDay {
				continue
			}
			yCount++
			if yCount == 1 {
				if trimmedFormat[i+1] != 'y' {
					yCount = 0
					fmt.Println("Error: invalid date format for year")
					os.Exit(1)
				}
				continue
			}

			if len(trimmedFormat)-1 == i || trimmedFormat[i+1] != 'y' {

				if yCount == 2 {
					formattedStr += "06"
					yCount = 0
				} else if yCount == 4 {
					formattedStr += "2006"
					yCount = 0
				} else {
					if trimmedFormat[i+1] != 'y' {
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

			//handle it for minute: mi
			if mCount == 1 && trimmedFormat[i+1] == 'i' {
				formattedStr += "04"
				mCount = 0
				continue
			}

			if len(trimmedFormat)-1 != i {

				if trimmedFormat[i+1] == 'o' {

					if len(trimmedFormat)-1 < i+3 {
						err := errors.New("Invalid date format")
						return " ", err
					}

					if trimmedFormat[i+2] == 'n' && trimmedFormat[i+3] == 't' && trimmedFormat[i+4] == 'h' {
						// isMonth = true
						formattedStr += "January"
						i += 4
						mCount = 0
						continue
					} else {
						err := errors.New("Invalid date format")
						return " ", err
					}
				}
				// isMonth = false
			}

			// if mCount == 1 {
			// 	if trimmedFormat[i+1] != 'm' {
			// 		formattedStr += "Jan"
			// 		mCount = 0
			// 		continue
			// 	}
			// }

			if mCount == 2 {
				// isMonth = false
				formattedStr += "Jan"
				mCount = 0
			}
		case 'd':
			// isMonth = false

			if len(trimmedFormat)-1 != i {

				if trimmedFormat[i+1] == 'a' {
					if trimmedFormat[i+2] == 'y' {
						formattedStr += "Mon"
						isDay = true
						continue
					}
				}
			}

			dCount++
			if dCount == 1 {
				if trimmedFormat[i+1] != 'd' {
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
		case 'h':

			hCount++
			if hCount == 1 {
				if trimmedFormat[i+1] != 'h' {
					err := errors.New("Error: Invalid format for time, hour must be in HH format")
					return "", err
				}
				continue
			}

			if hCount == 2 {
				//for 24 hour format
				if trimmedFormat[i+1] == '2' && trimmedFormat[i+2] == '4' {
					formattedStr += "15"
				} else {
					formattedStr += "03"
				}
			}

		case 's':
			sCount++
			if sCount == 1 {
				if trimmedFormat[i+1] != 's' {
					err := errors.New("Invalid time format, vlaid time format is hh:mi:ss")
					return "", err
				}
			}

			if sCount == 2 {
				formattedStr += "05"
			}

		case ':':
			formattedStr += ":"
		case 'p':
			if trimmedFormat[i+1] == '.' {
				i = i + 2
				formattedStr += "P.M"
			} else {
				i = i + 1
				formattedStr += "PM"
			}
		case '.':
			formattedStr += "."
		}
	}

	var formattedDate string

	if dateInt > 0 {

	} else {
		formattedDate = t.Format(formattedStr)
	}

	return formattedDate, nil

}
