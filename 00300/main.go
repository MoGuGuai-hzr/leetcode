package main

import "sort"

func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, v := range nums {
		if i := sort.SearchInts(dp, v); i < len(dp) {
			dp[i] = v
		} else {
			dp = append(dp, v)
		}
	}
	return len(dp)
}

func lengthOfLIS2(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		longer := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && longer < dp[j] {
				longer = dp[j]
			}
		}
		dp[i] = longer + 1
	}

	longest := 0
	for _, long := range dp {
		if longest < long {
			longest = long
		}
	}
	return longest
}
