package main

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	var cnt int
	for i := 2; i < n; i++ {
		diff := nums[i-1] - nums[i-2]
		for j := 0; i < n && nums[i]-nums[i-1] == diff; {
			i++
			j++
			cnt += j
		}
	}

	return cnt
}
