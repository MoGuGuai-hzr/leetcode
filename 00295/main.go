package main

import (
	"container/heap"
	"sort"
)

type MedianFinder struct {
	// 一个最小堆一个最大堆, 分别记录小于等于中位数的数, 和大于中位数的数
	// 最大堆里存负数
	queMin, queMax hp
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	minQ, maxQ := &mf.queMin, &mf.queMax
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {
		heap.Push(minQ, -num)
		if maxQ.Len()+1 < minQ.Len() {
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		heap.Push(maxQ, num)
		if maxQ.Len() > minQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	minQ, maxQ := mf.queMin, mf.queMax
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	}
	return float64(maxQ.IntSlice[0]-minQ.IntSlice[0]) / 2
}

// 使用(最小堆)堆实现优先队列
type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

// type MedianFinder struct {
//     nums        *redblacktree.Tree
//     total       int
//     left, right iterator
// }

// func Constructor() MedianFinder {
//     return MedianFinder{nums: redblacktree.NewWithIntComparator()}
// }

// func (mf *MedianFinder) AddNum(num int) {
//     if count, has := mf.nums.Get(num); has {
//         mf.nums.Put(num, count.(int)+1)
//     } else {
//         mf.nums.Put(num, 1)
//     }
//     if mf.total == 0 {
//         it := mf.nums.Iterator()
//         it.Next()
//         mf.left = iterator{it, 1}
//         mf.right = mf.left
//     } else if mf.total%2 == 1 {
//         if num < mf.left.Key().(int) {
//             mf.left.prev()
//         } else {
//             mf.right.next()
//         }
//     } else {
//         if mf.left.Key().(int) < num && num < mf.right.Key().(int) {
//             mf.left.next()
//             mf.right.prev()
//         } else if num >= mf.right.Key().(int) {
//             mf.left.next()
//         } else {
//             mf.right.prev()
//             mf.left = mf.right
//         }
//     }
//     mf.total++
// }

// func (mf *MedianFinder) FindMedian() float64 {
//     return float64(mf.left.Key().(int)+mf.right.Key().(int)) / 2
// }

// type iterator struct {
//     redblacktree.Iterator
//     count int
// }

// func (it *iterator) prev() {
//     if it.count > 1 {
//         it.count--
//     } else {
//         it.Prev()
//         it.count = it.Value().(int)
//     }
// }

// func (it *iterator) next() {
//     if it.count < it.Value().(int) {
//         it.count++
//     } else {
//         it.Next()
//         it.count = 1
//     }
// }

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
