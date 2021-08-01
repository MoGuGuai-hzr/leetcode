package main

func titleToNumber(columnTitle string) int {
	ans := 0
	base := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		ans += int((columnTitle[i] - 'A' + 1)) * base
		base *= 26
	}

	return ans
}
