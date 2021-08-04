package main

import (
	"fmt"
	"sort"
)

func triangleNumber1(nums []int) (ans int) {
	sort.Ints(nums)
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			ans += sort.SearchInts(nums[j+1:], v+nums[j])
		}
	}
	return
}

func triangleNumber(nums []int) (ans int) {
	n := len(nums)
	sort.Ints(nums)
	for i, v := range nums {
		k := i + 1
		for j := i + 1; j < n; j++ {
			for k+1 < n && nums[k+1] < v+nums[j] {
				k++
			}
			ans += max(k-j, 0)
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	for _, test := range []struct {
		nums []int
		exp  int
	}{
		{[]int{2, 2, 3, 4}, 3},
	} {
		r := triangleNumber(test.nums)
		fmt.Println(r, r == test.exp)
	}
}
