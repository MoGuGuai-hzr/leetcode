package main

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	list := make([]int, n+1)
	list[1] = 1
	max := 1
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			list[i] = list[i/2]
		} else {
			list[i] = list[i/2] + list[i/2+1]
		}
		if list[i] > max {
			max = list[i]
		}
	}

	return max
}

func main() {
	getMaximumGenerated(15)
}
