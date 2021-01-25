package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	result := make([][]int, numRows)
	result[0] = []int{1}

	for i := 1; i < numRows; i++ {
		result[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 {
				result[i][j] = result[i-1][j]
			} else if j == i {
				result[i][j] = result[i-1][j-1]
			} else {
				result[i][j] = result[i-1][j-1] + result[i-1][j]
			}
		}
	}

	return result
}
