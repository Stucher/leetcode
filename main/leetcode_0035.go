package main

func searchInsert(nums []int, target int) int {

	if target <= nums[0] {
		return 0
	}
	length := len(nums)
	if target > nums[length-1] {
		return length
	}
	index := 0
	for index = 1; index < length-1; index++ {
		if nums[index-1] < target && nums[index] >= target {
			break
		}
	}

	return index
}
