package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	y, tmpX := 0, x
	for tmpX > 0 {
		y = y*10 + tmpX%10
		tmpX /= 10
	}
	return x == y
}
