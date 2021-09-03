package main

import (
	"fmt"
	"math/rand"
)

func smallestK(arr []int, k int) []int {
	if len(arr) == 0 {
		return nil
	}
	left, right := 0, len(arr)
	for {
		r := random(left, right)
		v := arr[r]
		arr[r], arr[left] = arr[left], arr[r]
		i, j := left, right-1
		for i < j {
			for i < j && arr[j] > v {
				j--
			}
			arr[i], arr[j] = arr[j], arr[i]

			for i < j && arr[i] <= v {
				i++
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		if i == k {
			return arr[:k]
		}
		if i+1 == k {
			return arr[:k]
		}
		if i < k {
			left = i + 1
		} else {
			right = i
		}
	}
}

func random(start, end int) int {
	return rand.Int()%(end-start) + start
}

func main() {
	for _, test := range []struct {
		arr []int
		k   int
	}{
		{[]int{1, 3, 5, 7, 2, 4, 6, 8}, 4},
		{[]int{8, 6, 1, 3, 8, 1, 7, 3, 0, 24, 2, 8}, 4},
	} {
		r := smallestK(test.arr, test.k)
		fmt.Println(r)
	}
}
