package main

import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)
	max := 0
	n := len(nums)
	for i := 0; i < n/2; i++ {
		if nums[i]+nums[n-i-1] > max {
			max = nums[i] + nums[n-i-1]
		}
	}
	return max
}

func main() {
	for _, test := range []struct {
		nums []int
		exp  int
	}{
		{[]int{3, 5, 2, 3}, 7},
		{[]int{3, 5, 4, 2, 4, 6}, 8},
	} {
		r := minPairSum(test.nums)
		fmt.Println(r, r == test.exp)
	}
}
