package main

func removeElement(nums []int, val int) int {
	length := 0
	numLength := len(nums)
	for i := 0; i < numLength; i++ {
		if nums[i] == val {
			continue
		}
		nums[length] = nums[i]
		length++
	}
	return length
}
