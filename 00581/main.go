package main

import (
	"math"
	"sort"
)

func findUnsortedSubarray1(nums []int) int {
	nums2 := make([]int, len(nums))
	for i, v := range nums {
		nums2[i] = v
	}
	sort.Ints(nums2)
	left := -1
	right := len(nums)
	for i := 0; i < len(nums) && nums[i] == nums2[i]; i++ {
		left++
	}
	for i := len(nums) - 1; i >= 0 && nums[i] == nums2[i]; i-- {
		right--
	}
	if right <= left {
		return 0
	}

	return right - left - 1
}

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	minn, maxn := math.MaxInt64, math.MinInt64
	left, right := -1, -1
	for i, num := range nums {
		if maxn > num {
			right = i
		} else {
			maxn = num
		}
		if minn < nums[n-i-1] {
			left = n - i - 1
		} else {
			minn = nums[n-i-1]
		}
	}
	if right == -1 {
		return 0
	}
	return right - left + 1
}
