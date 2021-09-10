package main

func chalkReplacer(chalk []int, k int) int {
	var i = 0
	sum := 0
	for _, v := range chalk {
		sum += v
	}
	k = k % sum
	for ; chalk[i] <= k; i++ {
		k -= chalk[i]
	}
	return i
}
