func bade(collection []int) bool {
	var mem [10]bool
	for _, v := range collection {
		if v == 0 {
			return true
		}
		if mem[v] {
			return true
		} else {
			mem[v] = true
		}
	}
	return false
}

func isValidSudoku(board [9][9]int) bool {
	for y := 0; y < 9; y++ {
		if bade(board[y][:]) {
			return false
		}
	}
	for y := 0; y < 9; y++ {
		var tmp []int
		for x := 0; x < 9; x++ {
			tmp = append(tmp, board[x][y])
		}
		if bade(tmp) {
			return false
		}
	}
	for s := 0; s < 9; s++ {
		var tmp []int
		for y := int(s/3) * 3; y < int(s/3)*3+3; y++ {
			for x := int(s%3) * 3; x < int(s%3)*3+3; x++ {
				tmp = append(tmp, board[x][y])
			}
		}
		if bade(tmp) {
			return false
		}
	}

	return true
}