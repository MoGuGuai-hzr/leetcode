package main

import "fmt"

func pathInZigZagTree(label int) []int {
	h := 0
	for i := label; i > 0; {
		i >>= 1
		h++
	}

	ans := make([]int, h)
	for i := h - 1; i >= 0; i-- {
		ans[i] = label
		label >>= 1
	}

	i := 0
	n := 1
	if h%2 == 1 {
		i = 1
		n <<= 1
	}
	for ; i < h-1; i += 2 {
		ans[i] = 3*n - 1 - ans[i]
		n <<= 2
	}

	return ans
}

func main() {
	for _, test := range []struct {
		label int
		exp   []int
	}{
		{14, []int{1, 3, 4, 14}},
		{4, []int{1, 3, 4}},
	} {
		r := pathInZigZagTree(test.label)
		fmt.Println(r)
	}
}
