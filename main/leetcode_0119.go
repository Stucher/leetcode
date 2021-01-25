package main

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	result := make([]int, rowIndex+1)
	result[0] = 1
	times := 1

	for times <= rowIndex {
		for i := times; i >= 0; i-- {
			if i == 0 || i == times {
				result[i] = 1
			} else if i >= times/2 {
				result[i] = result[i] + result[i-1]
			} else {
				result[i] = result[times-i]
			}
		}
		times++
	}

	return result
}
