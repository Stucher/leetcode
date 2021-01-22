package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	index := -1

	rNeedle := []rune(needle)
	rHayStack := []rune(haystack)

	lNeedle := len(rNeedle)
	lHayStack := len(rHayStack)

	fmt.Println(rNeedle)
	fmt.Println(lNeedle)
	fmt.Println(rHayStack)
	fmt.Println(lHayStack)

	for i := 0; i < lHayStack; i++ {
		tmpI := i
		fmt.Printf("%d\n",tmpI)
		for j := 0; j < lNeedle; j++ {
			fmt.Printf("%d %d %c %c %d\n", i, j, rHayStack[i], rNeedle[j], tmpI)
			if tmpI >= lHayStack {
				return -1
			}
			if rNeedle[j] != rHayStack[tmpI] {
				break
			}
			tmpI++
			if j == lNeedle-1 {
				return i
			}
		}
	}

	return index
}
