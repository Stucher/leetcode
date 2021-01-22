package main

import "fmt"

func mySqrt(x int) int {
	return div2Pow(1, x, x)
}

func div2Pow(start, end, target int) int {
	mid := (end + start) / 2

	fmt.Println([]int{start, end, mid, target})

	if mid*mid > target {
		return div2Pow(start, mid, target)
	} else if mid*mid < target {
		if (mid+1)*(mid+1) > target {
			return mid
		}
		return div2Pow(mid, end, target)
	}

	return mid
}
