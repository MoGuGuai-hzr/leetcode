package main

import (
	"fmt"
	"sort"
)

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff > 1 || diff < -1 {
			arr[i] = arr[i-1] + 1
		}
	}

	if arr[len(arr)-1] > len(arr) {
		return len(arr)
	}
	return arr[len(arr)-1]
}

func main() {
	for _, test := range []struct {
		arr []int
		exp int
	}{
		{[]int{2, 2, 1, 2, 1}, 2},
		{[]int{1, 100, 1000}, 3},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{73, 98, 9}, 3},
	} {
		r := maximumElementAfterDecrementingAndRearranging(test.arr)
		fmt.Println(r, r == test.exp)
	}
}
