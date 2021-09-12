package main

func checkValidString(s string) bool {
	maxCnt := 0
	minCnt := 0
	for _, c := range s {
		switch c {
		case '(':
			minCnt++
			maxCnt++
		case ')':
			minCnt = max(minCnt-1, 0)
			maxCnt--
		case '*':
			minCnt = max(minCnt-1, 0)
			maxCnt++
		}
		if maxCnt < 0 {
			return false
		}
	}

	return minCnt == 0
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
