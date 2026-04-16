func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
    up := 0
    down := m - 1
    for up <= down {
        mid := up + (down - up) / 2
        row := matrix[mid]
        if target < row[0] {
            down = mid - 1
        } else if row[n - 1] < target {
            up = mid + 1
        } else {
            left := 0
            right := n - 1
            for left <= right {
                mid := left + (right - left) / 2
                if target < row[mid] {
                    right = mid - 1
                } else if row[mid] < target {
                    left = mid + 1
                } else {
                    return true
                }
            }
            return false
        }
    }

    return false
}
