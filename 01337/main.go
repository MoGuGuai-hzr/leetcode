package main

import (
	"container/heap"
	"fmt"
	"reflect"
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	type List struct {
		val   int
		index int
	}

	list := []List{}
	m := len(mat)
	n := len(mat[0])
	for i, vv := range mat {
		v := sort.Search(n, func(j int) bool {
			return vv[j] <= 0
		})
		list = append(list, List{v, i})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].val < list[j].val || list[i].val == list[j].val && list[i].index < list[j].index
	})

	rst := make([]int, 0, k)
	for i := 0; i < m && i < k; i++ {
		rst = append(rst, list[i].index)
	}
	return rst
}

func main() {
	type args struct {
		mat [][]int
		k   int
	}
	tests := []struct {
		args args
		want []int
	}{

		{args: args{[][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}, 3}, want: []int{2, 0, 3}},
		{args: args{[][]int{{1, 1, 0}, {1, 1, 0}, {1, 1, 1}, {1, 1, 1}, {0, 0, 0}, {1, 1, 1}, {1, 0, 0}}, 6}, want: []int{4, 6, 0, 1, 2, 3}},
	}
	for _, tt := range tests {
		if got := kWeakestRows(tt.args.mat, tt.args.k); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("kWeakestRows() = %v, want %v\n", got, tt.want)
		}
	}
}

// 官方题解, 用堆而不是排序
func kWeakestRows2(mat [][]int, k int) []int {
	h := hp{}
	for i, row := range mat {
		pow := sort.Search(len(row), func(j int) bool { return row[j] == 0 })
		h = append(h, pair{pow, i})
	}
	heap.Init(&h)
	ans := make([]int, k)
	for i := range ans {
		ans[i] = heap.Pop(&h).(pair).idx
	}
	return ans
}

type pair struct{ pow, idx int }
type hp []pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.pow < b.pow || a.pow == b.pow && a.idx < b.idx
}
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
