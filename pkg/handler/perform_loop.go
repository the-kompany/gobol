package handler

import (
	"strconv"
	"strings"
)

func (d *Data) PerformLoopBlock(tokens []string) {

	if strings.ToLower(tokens[2]) == "times" {
		times, _ := strconv.Atoi(tokens[1])

		for i := 0; i < times; i++ {
			for i := 3; i < len(tokens)-1; i++ {

				trimmed := strings.TrimSpace(tokens[i])

				switch strings.ToLower(trimmed) {

				case "display":
					var actionStr string

					if strings.HasPrefix(strings.TrimSpace(tokens[i+1]), "\"") {

						actionStr = trimmed + " " + tokens[i+1]

					} else {
						actionStr = trimmed + " " + tokens[i+1]

					}

					d.Display(actionStr)

				case "move":
					actionStr := trimmed + " " + tokens[i+1] + " " + tokens[i+2] + " " + tokens[i+3]
					d.Move(actionStr)
				}
			}

		}
	}

}
