package main

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	diffs := make([]int, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		diffs[i] = nums[i+1] - nums[i]
	}

	cnt := 1
	for i := 0; i < len(diffs); i++ {
		if diffs[i] == 0 && i < len(diffs)-1 {
			continue
		}
		cnt_j := 1
		for j, k_j, diff_j := i-1, 0, 0; j >= 0; j-- {
			diff_j += diffs[j]
			k_j += diff_j
			if k_j > k {
				break
			}
			cnt_j++
		}
		if cnt_j > cnt {
			cnt = cnt_j
		}
	}
	return cnt
}

func maxFrequency2(nums []int, k int) int {
	sort.Ints(nums)
	ans := 1
	for l, r, total := 0, 1, 0; r < len(nums); r++ {
		total += (nums[r] - nums[r-1]) * (r - l)
		for total > k {
			total -= nums[r] - nums[l]
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	for _, test := range []struct {
		nums   []int
		k, exp int
	}{
		{[]int{1, 2, 4}, 5, 3},
		{[]int{1, 4, 8, 13}, 5, 2},
		{[]int{3, 9, 6}, 2, 1},
		{[]int{1, 1, 1, 1, 1, 1, 1}, 5, 7},
		{[]int{1, 2, 3, 10}, 23, 3},
		{[]int{1, 2, 3, 10}, 24, 4},
		{[]int{9940, 9995, 9944, 9937, 9941, 9952, 9907, 9952, 9987, 9964, 9940, 9914, 9941, 9933, 9912, 9934, 9980, 9907, 9980, 9944, 9910, 9997}, 7925, 22},
	} {
		r := maxFrequency(test.nums, test.k)
		fmt.Println(r, r == test.exp)
	}
}
