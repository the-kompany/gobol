package utils

import "strings"

func Split2(val string) []string {
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

func IsSpace(r rune) bool {
	if r == ' ' || r == '\t' {
		return true
	}

	return false
}

//not used
func Split(val string) []string {
	var isStrStart bool
	var str string
	startPos := 0
	var identifier string
	var isAlphabet bool
	var parenthesisStart bool
	var singleQuoteStart bool

	strTokens := []string{}

	for k, v := range val {

		if parenthesisStart {
			if v == ')' {
				parenthesisStart = false
				identifier = val[startPos : k+1]
				strTokens = append(strTokens, identifier)
				startPos = k + 1
				isAlphabet = false
				continue
			}
			continue
		}

		if isStrStart {

			if v == '"' {
				str = val[startPos : k+1]
				isStrStart = false
				startPos = k + 1
				strTokens = append(strTokens, str)
				continue
			}

			continue
		}

		if singleQuoteStart {

			if v == '\'' {
				str = val[startPos : k+1]
				singleQuoteStart = false
				startPos = k + 1
				strTokens = append(strTokens, str)
				continue
			}

			continue
		}

		if v == '"' {

			startPos = k
			isStrStart = true
			continue
		}

		if v == '\'' {
			startPos = k
			singleQuoteStart = true
			continue
		}

		if !IsSpace(v) {
			isAlphabet = true
		}

		if IsSpace(v) {
			if !isAlphabet {
				startPos++
				continue
			}
			identifier = val[startPos:k]
			startPos = k + 1
			strTokens = append(strTokens, identifier)
			isAlphabet = false
		} else if k == len(val)-1 {

			identifier = val[startPos : k+1]
			strTokens = append(strTokens, identifier)

		} else if v == '(' {
			parenthesisStart = true
		}

	}

	return strTokens
}
