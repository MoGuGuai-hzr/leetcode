package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type pair struct{ right, height int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].height > h[j].height }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func getSkyline(buildings [][]int) (ans [][]int) {
	n := len(buildings)
	boundaries := make([]int, 0, n*2)
	for _, building := range buildings {
		boundaries = append(boundaries, building[0], building[1])
	}
	sort.Ints(boundaries)

	idx := 0
	h := hp{}
	for _, boundary := range boundaries {
		for idx < n && buildings[idx][0] <= boundary {
			heap.Push(&h, pair{buildings[idx][1], buildings[idx][2]})
			idx++
		}
		for len(h) > 0 && h[0].right <= boundary {
			heap.Pop(&h)
		}

		maxn := 0
		if len(h) > 0 {
			maxn = h[0].height
		}
		if len(ans) == 0 || maxn != ans[len(ans)-1][1] {
			ans = append(ans, []int{boundary, maxn})
		}
	}
	return
}

func getSkyline2(buildings [][]int) [][]int {
	const (
		Left   = 0
		Right  = 1
		Height = 2
		X      = 0
		Y      = 1
	)

	exp := [][]int{{buildings[0][Left], buildings[0][Height]}, {buildings[0][Right], Height}}

	var l, r, n int
	for index := 1; index < len(buildings); index++ {
		building := buildings[index]
		n = len(exp)

		// 不相交
		if building[Left] > exp[n-1][X] {
			exp = append(exp, []int{building[Left], building[Height]})
			exp = append(exp, []int{building[Right], 0})
		} else if building[Left] == exp[n-1][X] {
			// 相切
			// 不等高
			if building[Height] != exp[n-2][Y] {
				exp[n-1][Y] = building[Height]
				exp = append(exp, []int{building[Right], 0})
			} else {
				exp[n-1][X] = building[Right]
			}
		} else {
			// 相交
			l = sort.Search(n, func(i int) bool {
				if exp[i][X] >= building[Left] {
					return true
				}
				return false
			})

			if l-1 >= 0 && exp[l-1][Y] < building[Height] {
				exp = append(exp[:l], append([][]int{{building[Left], building[Height]}}, exp[l:]...)...)
				l += 1
			}
			n = len(exp)

			r = sort.Search(n, func(i int) bool {
				if exp[i][X] >= building[Right] {
					return true
				}
				return false
			})

			if r >= n {
				exp = append(exp, []int{building[Right], 0})
			}

			n = len(exp)

			cnt := 0
			for i := l; i < r; i++ {
				t := i + cnt
				if t > n {
					t = n
				}
				copy(exp[i-cnt:i], exp[i:t])
				if exp[i][Y] <= building[Height] {
					cnt++
					if i-1 >= 0 && exp[i-1][Y] > building[Height] {
						cnt--
						exp[l][Y] = building[Height]
					}
				}
			}
			exp = exp[:n-cnt]

		}
	}
	return exp
}

func equal(rst, exp [][]int) bool {
	if len(rst) != len(exp) {
		return false
	}
	for i := 0; i < len(rst); i++ {
		if rst[i][0] != exp[i][0] || rst[i][1] != exp[i][1] {
			return false
		}
	}
	return true
}

func main() {
	for _, test := range []struct {
		buildutils, skyline [][]int
	}{
		{[][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}, [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}}},
		{[][]int{{0, 2, 3}, {2, 5, 3}}, [][]int{{0, 3}, {5, 0}}},
		{[][]int{{2, 9, 10}, {9, 12, 15}}, [][]int{{2, 10}, {9, 15}, {12, 0}}},
		{[][]int{{1, 2, 1}, {1, 2, 2}, {1, 2, 3}}, [][]int{{2, 10}, {9, 15}, {12, 0}}},
	} {
		r := getSkyline2(test.buildutils)
		fmt.Println(r, equal(r, test.skyline))
	}
}
