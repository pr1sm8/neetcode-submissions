func searchMatrix(matrix [][]int, target int) bool {
	// 0 1 2 .. n
	// :
	// m
	m := len(matrix) - 1
	n := len(matrix[0]) - 1
	start := 0
	end := m
	for start != end {
		mid := (start + end + 1) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			start = mid
		} else {
			end = mid - 1
		}
	}
	row := start
	start = 0
	end = n
	// fmt.Println(row)
	// return false
	for start <= end {
		mid := (start + end) / 2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			start = mid+1
		} else {
			end = mid-1
		}
	}
	return false
}
