package main

import (
	"fmt"
	"sort"
)

type sorted []int

func (s *sorted) Len() int {
	return len(*s)
}

func (s *sorted) Less(i int, j int) bool {
	return (*s)[i] > (*s)[j]
}

func (s *sorted) Swap(i int, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func hIndex(citations []int) int {
	s := sorted(citations)
	sort.Sort(&s)
	c := 0
	n := len(citations)
	if citations[n-1] >= n {
		return n
	}
	for i := 1; i < n; i++ {
		if citations[i] <= i && citations[i-1] >= i {
			c = i
		}
	}
	return c
}

func hIndex2(citations []int) (h int) {
	n := len(citations)
	counter := make([]int, n+1)
	for _, citation := range citations {
		if citation >= n {
			counter[n]++
		} else {
			counter[citation]++
		}
	}
	for i, tot := n, 0; i >= 0; i-- {
		tot += counter[i]
		if tot >= i {
			return i
		}
	}
	return 0
}

func main() {
	for _, test := range []struct {
		citations []int
		exp       int
	}{
		{[]int{3, 0, 6, 1, 5}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 10}, 4},
		{[]int{1, 10}, 1},
		{[]int{10, 10}, 2},
	} {
		r := hIndex(test.citations)
		r2 := hIndex2(test.citations)
		fmt.Println(r, r == test.exp, r2 == test.exp)
	}
}
