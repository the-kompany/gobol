package handler

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func DateToStr(d *Data, date, inputFormat, format string) (string, error) {

	var t time.Time
	var err error
	var dateInt int64

	inputFormat = strings.TrimSpace(inputFormat)

	if !strings.HasPrefix(strings.TrimSpace(date), "\"") {
		if v, ok := d.Vars[date]; !ok {
			fmt.Println("Error: undefined variable ", date, "at line", d.Line)
			os.Exit(1)
		} else {
			date = v.(string)
		}

	}

	if len(inputFormat) < 1 {
		err := errors.New("Input format is required")
		return "", err
	}

	if len(inputFormat) > 0 {

		if strings.HasPrefix(inputFormat, "\"") {

			inputFormat = inputFormat[1 : len(inputFormat)-1]
		}

		inputFormat = strings.ToLower(inputFormat)

		dateInput, err := getDateLayout(inputFormat)

		if err != nil {
			err = errors.New("Inavlid input date format: " + err.Error())
			return "", err
		}

		if strings.HasPrefix(date, "\"") {
			date = date[1 : len(date)-1]
		}

		t, err = time.Parse(dateInput, date)

		if err != nil {
			err = errors.New("Inavlid date format: " + err.Error())
			return "", err
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

	var formattedDate string

	formattedStr, err := getDateLayout(trimmedFormat)

	if err != nil {
		return "", err
	}

	if dateInt > 0 {

	} else {
		formattedDate = t.Format(formattedStr)
	}

	return formattedDate, nil

}

func getDateLayout(dateStr string) (string, error) {

	yCount := 0
	mCount := 0
	dCount := 0
	hCount := 0
	sCount := 0
	formattedStr := ""

	isDay := false
	// isMonth := false

	for i := 0; i < len(dateStr); i++ {
		switch dateStr[i] {
		case 'y':
			if isDay {
				continue
			}
			yCount++
			if yCount == 1 {
				if dateStr[i+1] != 'y' {
					yCount = 0
					fmt.Println("Error: invalid date format for year")
					os.Exit(1)
				}
				continue
			}

			if len(dateStr)-1 == i || dateStr[i+1] != 'y' {

				if yCount == 2 {
					formattedStr += "06"
					yCount = 0
				} else if yCount == 4 {
					formattedStr += "2006"
					yCount = 0
				} else {
					if dateStr[i+1] != 'y' {
						yCount = 0
						fmt.Println("Error: invalid date format for year")
						os.Exit(1)
					}
				}
			}

		case '-':
			formattedStr += "-"
		case '/':
			formattedStr += "/"
		case 'm':
			isDay = false
			mCount++

			//handle it for minute: mi
			if mCount == 1 && dateStr[i+1] == 'i' {
				formattedStr += "04"
				mCount = 0
				continue
			}

			if len(dateStr)-1 != i {

				if dateStr[i+1] == 'o' {

					if len(dateStr)-1 < i+3 {
						err := errors.New("Invalid date format")
						return " ", err
					}

					if dateStr[i+2] == 'n' && dateStr[i+3] == 't' && dateStr[i+4] == 'h' {
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
				formattedStr += "01"
				mCount = 0
			}
		case 'd':
			// isMonth = false

			if len(dateStr)-1 != i {

				if dateStr[i+1] == 'a' {
					if dateStr[i+2] == 'y' {
						formattedStr += "Mon"
						isDay = true
						continue
					}
				}
			}

			dCount++
			if dCount == 1 {
				if dateStr[i+1] != 'd' {
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
				if dateStr[i+1] != 'h' {
					err := errors.New("Error: Invalid format for time, hour must be in HH format")
					return "", err
				}
				continue
			}

			if hCount == 2 {
				//for 24 hour format
				if dateStr[i+1] == '2' && dateStr[i+2] == '4' {
					formattedStr += "15"
				} else {
					formattedStr += "03"
				}
			}

		case 's':
			sCount++
			if sCount == 1 {
				if dateStr[i+1] != 's' {
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
			if dateStr[i+1] == '.' {
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

	return formattedStr, nil

}

func (d *Data) CurrentDate() string {

	t := time.Now()

	formattedStr := t.Format("01-02-2006")

	return formattedStr

}

func (d *Data) CurrentTime() string {

	t := time.Now()

	formattedStr := t.Format("03:04:05pm")

	return formattedStr

}

func (d *Data) CurrentDateTime() string {
	t := time.Now()

	formattedStr := t.Format("01-02-2006 03:04:05pm")

	return formattedStr
}
