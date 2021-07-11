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
		fmt.Println(r, r == test.exp)
	}
}
