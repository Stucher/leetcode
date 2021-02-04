package main

func majorityElement(nums []int) int {
	length := len(nums)
	num, count := nums[0], 1

	for index := 1; index < length; index++ {
		if count == 0 {
			num, count = nums[index], 1
			continue
		}

		if nums[index] == num {
			count++
		} else {
			count--
		}
	}

	return num
}
