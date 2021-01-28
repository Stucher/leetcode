package main

import (
	"strings"
)

func isPalindrome2(s string) bool {
	if len(s) == 0 {
		return true
	}

	runeSlice := []rune(strings.ToLower(s))

	for lP, rP := 0, len(runeSlice)-1; lP < rP; {
		if !isLetter(runeSlice[lP]) {
			lP++
			continue
		}
		if !isLetter(runeSlice[rP]) {
			rP--
			continue
		}

		if runeSlice[lP] != runeSlice[rP] {
			return false
		}
		lP++
		rP--
	}

	return true
}

func isLetter(charInt rune) bool {
	indexA, indexZ, indexLowerA, indexLowerZ, zero, nine := rune('A'), rune('Z'), rune('a'), rune('z'), rune('0'), rune('9')
	return (charInt >= indexA && charInt <= indexZ) || (charInt >= indexLowerA && charInt <= indexLowerZ) || (charInt >= zero && charInt <= nine)
}
