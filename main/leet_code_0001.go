package main

import (
	"sort"
)

func twoSum(nums []int, target int) []int {
	tmpNums := make([]int, len(nums))
	copy(tmpNums, nums)
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	tmpRes := []int{-1, -1}
	for i < j {
		if nums[i]+nums[j] < target {
			i++
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			tmpRes[0] = nums[i]
			tmpRes[1] = nums[j]
			break
		}
	}

	res := []int{-1, -1}
	for index, value := range tmpNums {
		if value == tmpRes[0] && res[0] == -1 {
			res[0] = index
			continue
		}
		if value == tmpRes[1] && res[1] == -1 {
			res[1] = index
			continue
		}
	}

	return res
}
