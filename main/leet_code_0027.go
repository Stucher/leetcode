package main

import (
	"math"
	"strings"
)

func addBinary(a string, b string) string {

	var deltaBit uint8 = 0
	maxLen := int(math.Max(float64(len(a)), float64(len(b)))) + 1
	result := make([]uint8, maxLen)
	i, j, k := len(a)-1, len(b)-1, maxLen-1

	// 逆序相加
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		bitResult := a[i] + b[j] - 2*uint8('0') + deltaBit
		if bitResult >= 2 {
			deltaBit = 1
			bitResult -= 2
		} else {
			deltaBit = 0
		}

		result[k] = bitResult + uint8('0')
		k--
	}

	// 处理剩余值
	for ; i >= 0; i-- {
		bitResult := a[i] - uint8('0') + deltaBit
		if bitResult >= 2 {
			deltaBit = 1
			bitResult -= 2
		} else {
			deltaBit = 0
		}

		result[k] = bitResult + uint8('0')
		k--
	}

	// 处理剩余值
	for ; j >= 0; j-- {
		bitResult := b[j] - uint8('0') + deltaBit
		if bitResult >= 2 {
			deltaBit = 1
			bitResult -= 2
		} else {
			deltaBit = 0
		}

		result[k] = bitResult + uint8('0')
		k--
	}

	if deltaBit == 1 {
		result[k] = 1 + uint8('0')
	}

	resultString := strings.TrimLeft(string(result), string(uint8(0))+"0")

	if resultString == "" {
		return "0"
	}

	return resultString
}
