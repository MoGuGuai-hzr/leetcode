package main

import "fmt"

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	t := len(nums) / 2
	for _, v := range nums {
		if m[v] > t {
			return v
		}
	}
	return -1
}

func BoyerMoore(nums []int) int {
	var candidate, cnt, reCnt int
	for _, v := range nums {
		if cnt == 0 {
			candidate = v
		}
		if candidate == v {
			cnt++
		} else {
			cnt--
		}
	}

	for _, v := range nums {
		if v == candidate {
			reCnt++
		}
	}

	if reCnt > len(nums)/2 {
		return candidate
	}

	return -1
}
func main() {
	for _, test := range []struct {
		nums []int
		exp  int
	}{
		{[]int{1, 2, 5, 9, 5, 9, 5, 5, 5}, 5},
		{[]int{3, 2}, -1},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	} {
		r := majorityElement(test.nums)
		r2 := BoyerMoore(test.nums)
		fmt.Println(r == test.exp, r2 == test.exp)
	}
}
