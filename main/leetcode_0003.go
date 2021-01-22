package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charMap := map[string]uint16{}
	// 遍历做map
	for _, ic := range s {
		c := fmt.Sprintf("%c", ic)
		cnt, ok := charMap[c]
		if ok {
			charMap[c] = cnt + 1
		} else {
			charMap[c] = 1
		}
	}
	// 缩小窗口
	return 0
}
