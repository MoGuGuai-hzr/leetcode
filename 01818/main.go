package main

import (
	"fmt"
	"sort"
)

func minAbsoluteSumDiff(nums1, nums2 []int) int {
	rec := append(sort.IntSlice(nil), nums1...)
	rec.Sort()
	sum, maxn, n := 0, 0, len(nums1)
	for i, v := range nums2 {
		diff := abs(nums1[i] - v)
		sum += diff
		j := rec.Search(v)
		// Search 返回的是 v 应该插入的位置(之后的数往后移)
		// 第一个检查判断右边的数
		// 第二个检查判断左边的数
		if j < n {
			maxn = max(maxn, diff-(rec[j]-v))
		}
		if j > 0 {
			maxn = max(maxn, diff-(v-rec[j-1]))
		}
	}
	return (sum - maxn) % (1e9 + 7)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	for _, test := range []struct {
		nums1, nums2 []int
		exp          int
	}{
		{[]int{1, 7, 5}, []int{2, 3, 5}, 3},
		{[]int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10}, 0},
		{[]int{1, 10, 4, 4, 2, 7}, []int{9, 3, 5, 1, 7, 4}, 20},
		{[]int{1, 28, 21}, []int{9, 21, 20}, 9},
	} {
		r := minAbsoluteSumDiff(test.nums1, test.nums2)
		fmt.Println(r, r == test.exp)
	}
}
