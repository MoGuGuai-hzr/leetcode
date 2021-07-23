package main

import "fmt"

func isCovered(ranges [][]int, left int, right int) bool {
	m := make([]bool, right-left+1, right-left+1)
	for _, ran := range ranges {
		if ran[0] >= left || ran[0] <= right {
			var i int
			if ran[0] >= left {
				i = ran[0] - left
			}
			for ; i <= ran[1]-left && i <= right-left; i++ {
				m[i] = true
			}
		}
	}
	fmt.Println(m)
	for _, b := range m {
		if !b {
			return false
		}
	}
	return true
}

func main() {
	for _, test := range []struct {
		ranges      [][]int
		left, right int
		exp         bool
	}{
		{[][]int{{1, 2}, {3, 4}, {5, 6}}, 2, 5, true},
		{[][]int{{1, 10}, {10, 20}}, 21, 21, false},
	} {
		r := isCovered(test.ranges, test.left, test.right)
		fmt.Println(r, r == test.exp)
	}
}
