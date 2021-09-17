package main

func isValidSudoku2(board [][]byte) bool {
	indexList := make([]int, 27)
	for i := 0; i < 27; i++ {
		indexList[i] = 1 << i
	}
	m := make([]int, 9)
	for rowIndex, row := range board {
		for columnIndex, s := range row {
			if s == '.' {
				continue
			}

			v := int(s - '1')
			// 在 rowIndex 行出现
			if m[v]&indexList[rowIndex] != 0 {
				return false
			}
			m[v] |= indexList[rowIndex]

			// 在 columnIndex 列出现
			if m[v]&indexList[9+columnIndex] != 0 {
				return false
			}
			m[v] |= indexList[9+columnIndex]

			// 在第 index 个方格中出现
			index := columnIndex/3 + 3*(rowIndex/3)
			if m[v]&indexList[18+index] != 0 {
				return false
			}
			m[v] |= indexList[18+index]
		}
	}

	return true
}

func isValidSudoku1(board [][]byte) bool {
	indexList := make([]int, 27)
	for i := 0; i < 27; i++ {
		indexList[i] = 1 << i
	}
	m := make(map[int]int, 9)
	for rowIndex, row := range board {
		for columnIndex, s := range row {
			if s == '.' {
				continue
			}

			v := int(s - '0')
			// 在 rowIndex 行出现
			if m[v]&indexList[rowIndex] != 0 {
				return false
			}
			m[v] |= indexList[rowIndex]

			// 在 columnIndex 列出现
			if m[v]&indexList[9+columnIndex] != 0 {
				return false
			}
			m[v] |= indexList[9+columnIndex]

			// 在第 index 个方格中出现
			index := columnIndex/3 + 3*(rowIndex/3)
			if m[v]&indexList[18+index] != 0 {
				return false
			}
			m[v] |= indexList[18+index]
		}
	}

	return true
}

func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]int
	var subboxes [3][3][9]int
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++
			columns[j][index]++
			subboxes[i/3][j/3][index]++
			if rows[i][index] > 1 || columns[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
