package solution

func maxMatrixSum(matrix [][]int) int64 {
	negativeCount := 0
	min := matrix[0][0]
	// minIsNegative := false
	if min < 0 {
		min = -min
		// minIsNegative = true
	}
	sum := 0
	for i := range matrix {
		for j := range matrix[i] {
			current := matrix[i][j]
			if matrix[i][j] < 0 {
				current = -current
				negativeCount++
			}
			// if matrix[i][j] == 0 {
			// 	zeroCount++
			// }
			if current < min {
				min = current
				if matrix[i][j] < 0 {
					// minIsNegative = true
				}
			}
			sum += current
		}
	}

	if negativeCount%2 == 0 {
		return int64(sum)
	}
	return int64(sum - 2*min)
	// return 0
}
