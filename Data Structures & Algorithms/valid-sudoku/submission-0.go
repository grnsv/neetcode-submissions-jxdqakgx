func isValidSudoku(board [][]byte) bool {
	seen := make([]bool, 9)
	clear := func() {
		for i := range seen {
			seen[i] = false
		}
	}

	for i := range 9 {
		for j := range 9 {
			if board[i][j] == '.' {
				continue
			}

			num := board[i][j] - '1'
			if seen[num] {
				return false
			}

			seen[num] = true
		}

		clear()
	}

	for j := range 9 {
		for i := range 9 {
			if board[i][j] == '.' {
				continue
			}

			num := board[i][j] - '1'
			if seen[num] {
				return false
			}

			seen[num] = true
		}

		clear()
	}

	for square := range 9 {
		for cell := range 9 {
			i := (square/3) * 3 + cell/3
			j := (square%3) * 3 + cell%3
			if board[i][j] == '.' {
				continue
			}

			num := board[i][j] - '1'
			if seen[num] {
				return false
			}

			seen[num] = true
		}

		clear()
	}

	return true
}
