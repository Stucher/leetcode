package main

func plusOne(digits []int) []int {

	// 反转数组
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	ret := make([]int, 0)
	ten := 0
	digits[0]++

	for i := 0; i < len(digits); i++ {
		ret = append(ret, (digits[i]+ten)%10)
		if digits[i]+ten >= 10 {
			ten = 1
		} else {
			ten = 0
		}
	}
	if ten > 0 {
		ret = append(ret, 1)
	}

	// 反转数组
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}

	return ret
}
