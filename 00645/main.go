package main

import "fmt"

func findErrorNums(nums []int) []int {
	r := []int{0, 0}
	m := make([]bool, len(nums), len(nums))
	for _, v := range nums {
		if m[v-1] {
			r[0] = v
		}
		m[v-1] = true
	}
	for i, b := range m {
		if !b {
			r[1] = i + 1
			break
		}
	}
	return r
}

func equal(r, e []int) bool {
	if len(r) != len(e) {
		return false
	}
	for i := 0; i < len(r); i++ {
		if r[i] != e[i] {
			return false
		}
	}
	return true
}

func main() {
	for _, test := range []struct {
		nums   []int
		expect []int
	}{
		{[]int{1, 2, 2, 4}, []int{2, 3}},
		{[]int{1, 1}, []int{1, 2}},
		{[]int{3, 2, 2}, []int{2, 1}},
	} {
		r := findErrorNums(test.nums)
		fmt.Println(r, test.expect, equal(r, test.expect))
	}
}
