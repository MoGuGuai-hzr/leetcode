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
	for _, b := range m {
		if !b {
			return false
		}
	}
	return true
}

func isCovered2(ranges [][]int, left int, right int) bool {
	n := right - left + 1 + 2
	m := make([]int, n, n)
	for _, ran := range ranges {
		if !(ran[0] > right || ran[1] < left) {
			if ran[0] < left {
				m[0]++
			} else {
				m[ran[0]-left+1]++
			}

			if ran[1] > right {
				m[n-1]--
			} else {
				m[ran[1]-left+1+1]--
			}
		}
	}
	cnt := m[0]
	for i := 1; i < n-1; i++ {
		cnt += m[i]
		if cnt <= 0 {
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
		{[][]int{{1, 50}}, 1, 50, true},
	} {
		r := isCovered(test.ranges, test.left, test.right)
		r2 := isCovered2(test.ranges, test.left, test.right)
		fmt.Println(r, r == test.exp, "r2:", r2, r2 == test.exp)
	}
}
