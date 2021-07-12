package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	h := 0
	n := len(citations)
	for i := n - 1; i >= 0 && citations[i] >= n-i; i-- {
		h = n - i
	}
	return h
}

func hIndex2(citations []int) int {
	n := len(citations)
	return n - sort.Search(n, func(x int) bool {
		return citations[x] >= n-x
	})
}

func main() {
	for _, test := range []struct {
		citations []int
		exp       int
	}{
		{[]int{0, 1, 3, 5, 6}, 3},
		{[]int{0, 0}, 0},
		{[]int{1, 2, 3, 4, 5, 6, 10}, 4},
		{[]int{1, 10}, 1},
		{[]int{10, 10}, 2},
	} {
		r := hIndex(test.citations)
		r2 := hIndex2(test.citations)
		fmt.Println(r, r == test.exp, r2 == test.exp)
	}
}
