package main

func merge(nums1 []int, m int, nums2 []int, n int) {

	// nums1的0~m、nums2的0~n 按从大到小 从后向前插到nums1里
	for p := m + n; m > 0 && n > 0; p-- {
		if nums2[n-1] >= nums1[m-1] {
			nums1[p-1] = nums2[n-1]
			n--
		} else {
			nums1[p-1] = nums1[m-1]
			m--
		}
	}

	// 补充剩余的nums2小的数据
	for ; n-1 >= 0; n-- {
		nums1[n-1] = nums2[n-1]
	}

}
