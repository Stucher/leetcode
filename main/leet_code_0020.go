package main

import (
	"fmt"
	"strings"
)

func cp(str string) uint8 {
	if str == ")" {
		return uint8('(')
	} else if str == "]" {
		return uint8('[')
	} else if str == "}" {
		return uint8('{')
	} else {
		return 0
	}
}

func isValid(s string) bool {

	if s == "" {
		return true
	}

	resStr := ""

	for _, cInt := range s {
		c := fmt.Sprintf("%c", cInt)
		switch c {
		case "(", "[", "{":
			resStr = strings.Join([]string{resStr, c}, "")
		case ")", "]", "}":
			if len(resStr) == 0 || resStr[len(resStr)-1] != cp(c) {
				return false
			} else {
				if len(resStr) == 1 {
					resStr = ""
				} else {
					resStr = resStr[0 : len(resStr)-1]
				}
			}
		}
	}

	return len(resStr) == 0
}
