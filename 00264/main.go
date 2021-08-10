package main

import (
	"container/heap"
	"math"
	"sort"
)

func nthSuperUglyNumber1(n int, primes []int) (ugly int) {
	seen := map[int]bool{1: true}
	h := &hp{[]int{1}}
	for i := 0; i < n; i++ {
		ugly = heap.Pop(h).(int)
		for _, prime := range primes {
			if next := ugly * prime; !seen[next] {
				seen[next] = true
				heap.Push(h, next)
			}
		}
	}
	return
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func nthSuperUglyNumber2(n int, primes []int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	m := len(primes)
	pointers := make([]int, m)
	for i := range pointers {
		pointers[i] = 1
	}
	for i := 2; i <= n; i++ {
		nums := make([]int, m)
		minNum := math.MaxInt64
		for j, p := range pointers {
			nums[j] = dp[p] * primes[j]
			minNum = min(minNum, nums[j])
		}
		dp[i] = minNum
		for j, num := range nums {
			if minNum == num {
				pointers[j]++
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
