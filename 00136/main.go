package main

import "fmt"

func singleNumber(nums []int) int {
	r := 0
	for i := 0; i < len(nums); i++ {
		r ^= nums[i]
	}
	return r
}

func main() {
	for _, test := range []struct {
		nums []int
		exp  int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	} {
		r := singleNumber(test.nums)
		fmt.Println(r, r == test.exp)
	}
}
