package main

import "fmt"

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	length := 0
	tmpNum := nums[0] - 1
	initLength := len(nums)

	for i := 0; i < initLength; i++ {
		if tmpNum == nums[i] {
			continue
		}
		tmpNum = nums[i]
		nums[length] = nums[i]
		length++
		fmt.Printf("%d ", tmpNum)
	}
	fmt.Println()

	return length
}
