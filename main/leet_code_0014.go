package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	result := ""

	strMinLen := len(strs[0])
	for _, tmpStr := range strs {
		if len(tmpStr) < strMinLen {
			strMinLen = len(tmpStr)
		}
	}

	for idx := 0; idx < strMinLen; idx++ {
		tmpC := strs[0][idx]
		canAppend := true
		for _, tmpStr := range strs {
			if tmpStr[idx] != tmpC {
				canAppend = false
				break
			}
		}
		if canAppend {
			result = strings.Join([]string{result, fmt.Sprintf("%c", tmpC)}, "")
		} else {
			break
		}
	}

	return result
}
