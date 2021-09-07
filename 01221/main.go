package main

func balancedStringSplit(s string) int {
	cnt := 0
	for i, r, l := 0, 0, 0; i < len(s); i++ {
		switch s[i] {
		case 'R':
			r++
		case 'L':
			l++
		}
		if l == r {
			cnt++
		}
	}
	return cnt
}
