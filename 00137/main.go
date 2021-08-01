package main

import "fmt"

func singleNumber(nums []int) int {
	r := nums[0]
	for i := 1; i < len(nums); i++ {
		r = ^(r ^ nums[i])
	}
	return r
}

func main() {
	for _, test := range []struct {
		nums []int
		exp  int
	}{
		{[]int{2, 2, 3, 2}, 3},
		{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
	} {
		r := singleNumber(test.nums)
		fmt.Println(r, r == test.exp)
	}

	var a int = 10
	r := a
	fmt.Println(r)
	r = ^(r ^ a)
	fmt.Println(r)
	r = ^(r ^ a)
	fmt.Println(r)

	fmt.Println(^(^(a ^ a) ^ a))
}
