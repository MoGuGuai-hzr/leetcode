package main

import (
	"fmt"
	"sort"
)

func search(nums []int, target int) int {
	n := len(nums)
	left := sort.Search(n, func(i int) bool {
		if nums[i] >= target {
			return true
		}
		return false
	})

	if left == n {
		return 0
	}

	right := sort.Search(n-left, func(i int) bool {
		if nums[left+i] > target {
			return true
		}
		return false
	})
	return right
}

func main() {
	for _, test := range []struct {
		nums        []int
		target, exp int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 4, 0},
		{[]int{5, 7, 7, 8, 8, 10}, 5, 1},
		{[]int{5, 7, 7, 8, 8, 10}, 6, 0},
		{[]int{5, 7, 7, 8, 8, 10}, 7, 2},
		{[]int{5, 7, 7, 8, 8, 10}, 8, 2},
		{[]int{5, 7, 7, 8, 8, 10}, 9, 0},
		{[]int{5, 7, 7, 8, 8, 10}, 10, 1},
		{[]int{5, 7, 7, 8, 8, 10}, 11, 0},
	} {
		r := search(test.nums, test.target)
		fmt.Println(r, r == test.exp)
	}
}
