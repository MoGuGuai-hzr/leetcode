package main

import "fmt"

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func main() {
	for _, test := range []struct {
		nums []int
		exp  int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	} {
		r := maxSubArray(test.nums)
		fmt.Println(r, r == test.exp)
	}
}
