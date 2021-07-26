package main

import (
	"fmt"
	"sort"
)

func minOperations(target, arr []int) int {
	n := len(target)
	pos := make(map[int]int, n)
	for i, val := range target {
		pos[val] = i
	}
	d := []int{}
	for _, val := range arr {
		if idx, has := pos[val]; has {
			if p := sort.SearchInts(d, idx); p < len(d) {
				d[p] = idx
			} else {
				d = append(d, idx)
			}
		}
	}
	return n - len(d)
}

// 垃圾算法
func minOperations2(target []int, arr []int) int {
	n := len(target)

	exist := make(map[int]struct{}, n)
	for _, v := range target {
		exist[v] = struct{}{}
	}

	m := make(map[int][]int, n)
	for i, v := range arr {
		if _, ok := exist[v]; ok {
			m[v] = append(m[v], i)
		}
	}

	stack := make([]int, 0, n)
	rst := n - minValue(m, target, 0, n, stack, 0)

	return rst
}

func minValue(m map[int][]int, target []int, i, n int, stack []int, max int) int {
	depth := len(stack)
	if i == n {
		return depth
	}

	list := m[target[i]]
	lenth := len(list)

	var index int
	if depth != 0 {
		index = sort.Search(lenth, func(i int) bool {
			if list[i] > stack[depth-1] {
				return true
			}
			return false
		})
	}

	if index < lenth {
		stack = append(stack, list[index])
		push := minValue(m, target, i+1, n, stack, depth+1)
		if push-depth >= n-i {
			return push
		}
		skip := minValue(m, target, i+1, n, stack[:depth], depth)
		if push > skip {
			return push
		}
		return skip
	}

	return minValue(m, target, i+1, n, stack[:depth], depth)
}

func main() {
	for _, test := range []struct {
		target, arr []int
		exp         int
	}{
		{[]int{5, 1, 3}, []int{9, 4, 2, 3, 4}, 2},
		{[]int{6, 4, 8, 1, 3, 2}, []int{4, 7, 6, 2, 3, 8, 6, 1}, 3},
		{[]int{16, 7, 20, 11, 15, 13, 10, 14, 6, 8}, []int{11, 14, 15, 7, 5, 5, 6, 10, 11, 6}, 6},
	} {
		r := minOperations(test.target, test.arr)
		r2 := minOperations2(test.target, test.arr)
		fmt.Println(r, r2, r == test.exp, r2 == test.exp)
	}
}
