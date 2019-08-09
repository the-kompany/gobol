package utils

import "strings"

func Split(val string) []string {
	// pos := 0
	splitted := strings.Split(val, " ")
	fields := []string{}

	var (
		ok               = false
		s                string
		startParenthesis = false
		parenthesisStr   string
	)

	for _, v := range splitted {

		if strings.Contains(v, "(") {
			startParenthesis = true
			parenthesisStr = v
			if strings.HasSuffix(v, ")") {
				fields = append(fields, v)
				startParenthesis = false
			}
			continue

		}

		if startParenthesis {
			parenthesisStr += " "
			if strings.HasSuffix(v, ")") {
				startParenthesis = false
				parenthesisStr += v
				fields = append(fields, parenthesisStr)
				continue
			}

			parenthesisStr += v

		}

		if !startParenthesis {

			if strings.HasPrefix(v, "\"") {

				if strings.HasSuffix(v, "\"") {
					fields = append(fields, v)
					continue
				}

				ok = true

				s += v

				continue
			}

			if ok {
				s += " "
				if strings.HasSuffix(v, "\"") {
					// end = true
					ok = false

					s += v
					fields = append(fields, s)
					continue
				}
				s += v
				continue

			}

			if v == "" {
				continue
			}
			fields = append(fields, v)
		}

	}

	return fields

}
