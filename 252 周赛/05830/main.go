package main

func isThree(n int) bool {
	cnt := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			cnt++
			if cnt > 3 {
				return false
			}
		}
	}
	if cnt == 3 {
		return true
	}
	return false
}
