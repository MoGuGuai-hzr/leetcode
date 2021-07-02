package main

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (h *minHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *minHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func maxIceCream(costs []int, coins int) int {
	mh := minHeap(costs)
	heap.Init(&mh)
	count := 0
	for len(mh) > 0 {
		x := heap.Pop(&mh).(int)
		coins -= x
		if coins < 0 {
			break
		}
		count++
	}
	return count
}

func main() {
	for _, test := range []struct {
		costs         []int
		coins, expect int
	}{
		{[]int{1, 3, 2, 4, 1}, 7, 4},
		{[]int{10, 6, 8, 7, 7, 8}, 5, 0},
		{[]int{1, 6, 3, 1, 2, 5}, 20, 6},
	} {
		r := maxIceCream(test.costs, test.coins)
		fmt.Println(r, test.expect)
	}
}
