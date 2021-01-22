package main

import "fmt"

func romanToInt(s string) int {
	romanIntMapping := make(map[string]int)

	romanIntMapping["I"] = 1
	romanIntMapping["V"] = 5
	romanIntMapping["X"] = 10
	romanIntMapping["L"] = 50
	romanIntMapping["C"] = 100
	romanIntMapping["D"] = 500
	romanIntMapping["M"] = 1000

	length := len(s)
	i := 0
	result := 0
	for i < length {
		c1, c2 := fmt.Sprintf("%c", s[i]), ""
		if i+1 < length {
			c2 = fmt.Sprintf("%c", s[i+1])
		}

		if romanIntMapping[c1] < romanIntMapping[c2] {
			result = result + romanIntMapping[c2] - romanIntMapping[c1];
			i += 2
		} else {
			result += romanIntMapping[c1]
			i += 1
		}
	}

	return result
}
