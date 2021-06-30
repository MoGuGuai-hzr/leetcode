package main

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"time"
)

type walked struct {
	bus, route map[int]bool
}

func (wRaw walked) clone() (w walked) {
	w = walked{bus: make(map[int]bool), route: make(map[int]bool)}
	for k, v := range wRaw.bus {
		w.bus[k] = v
	}
	for k, v := range wRaw.route {
		w.route[k] = v
	}
	return
}

func (w walked) allBus() string {
	buses := make([]int, 0, len(w.bus))
	for bus := range w.bus {
		buses = append(buses, bus)
	}
	sort.Ints(buses)
	return fmt.Sprint(buses)
}

func InSlice(list []int, x int) bool {
	for _, v := range list {
		if v == x {
			return true
		}
	}
	return false
}

func numBusesToDestination1(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	list := list.New()
	set := make(map[string]bool)

	for id, route := range routes {
		if InSlice(route, source) {
			if InSlice(route, target) {
				return 1
			} else {
				w := walked{bus: make(map[int]bool), route: make(map[int]bool)}
				w.bus[id] = true
				for _, station := range route {
					w.route[station] = true
				}
				list.PushBack(w)
				set[w.allBus()] = true
			}
		}
	}

	for i := 1; i < len(routes); i++ {
		fBreak := false
		el := list.Front()
		end := list.Back()
		for {
			if el == nil {
				break
			}
			for id, route := range routes {
				w := el.Value.(walked).clone()
				if w.bus[id] {
					continue
				}
				isConnected := false
				fEnd := false
				for _, station := range route {
					if w.route[station] {
						isConnected = true
					}
					if target == station {
						fEnd = true
					}
				}
				if isConnected && fEnd {
					return len(w.bus) + 1
				}
				if isConnected {
					w.bus[id] = true
					for _, station := range route {
						w.route[station] = true
					}
					if !set[w.allBus()] {
						list.PushBack(w)
					}
					set[w.allBus()] = true
				}
			}
			lastEl := el
			el = el.Next()
			if lastEl == end {
				fBreak = true
			}
			list.Remove(lastEl)
			if fBreak {
				break
			}
		}
	}

	return -1
}

// 官网的解答, 速度差了一千倍, 太菜了...
// https://leetcode-cn.com/problems/bus-routes/solution/gong-jiao-lu-xian-by-leetcode-solution-yifz/
func numBusesToDestination(routes [][]int, source, target int) int {
	if source == target {
		return 0
	}

	// 基本思路是先建图, 在广度搜索
	// 图中点的数量
	n := len(routes)
	// 利用矩阵存边, 变的权重为1
	edge := make([][]bool, n)
	for i := range edge {
		edge[i] = make([]bool, n)
	}

	// 建图
	// rec 记录一个站台链接到哪些公交
	rec := map[int][]int{}
	for i, route := range routes {
		for _, site := range route {
			for _, j := range rec[site] {
				edge[i][j] = true
				edge[j][i] = true
			}
			rec[site] = append(rec[site], i)
		}
	}

	// 记录从起点公交到该公交的距离, -1 表示到不了(初始值)
	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	// 找到经过起点的公交, 使用切片模拟队列
	q := []int{}
	for _, bus := range rec[source] {
		dis[bus] = 1
		q = append(q, bus)
	}
	// 这里算出了起点到所有站台的距离
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		// 由于只考虑转乘次数, 所以前面使用过该公交, 后面就不用去了; 即使再使用, 次数也不会下降
		for y, b := range edge[x] {
			if b && dis[y] == -1 {
				dis[y] = dis[x] + 1
				q = append(q, y)
			}
		}
	}

	ans := math.MaxInt32
	for _, bus := range rec[target] {
		if dis[bus] != -1 && dis[bus] < ans {
			ans = dis[bus]
		}
	}
	if ans < math.MaxInt32 {
		return ans
	}
	return -1
}

func main() {
	for _, test := range []struct {
		routes                 [][]int
		source, target, expect int
	}{
		{[][]int{{1, 7}, {3, 5}}, 5, 5, 0},
		{[][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6, 2},
		{[][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12, -1},
		{[][]int{{25, 33}, {3, 5, 13, 22, 23, 29, 37, 45, 49}, {15, 16, 41, 47}, {5, 11, 17, 23, 33}, {10, 11, 12, 29, 30, 39, 45}, {2, 5, 23, 24, 33}, {1, 2, 9, 19, 20, 21, 23, 32, 34, 44}, {7, 18, 23, 24}, {1, 2, 7, 27, 36, 44}, {7, 14, 33}}, 7, 47, -1},
	} {
		start := time.Now()
		result := numBusesToDestination(test.routes, test.source, test.target)
		fmt.Println("官网程序", result, test.expect, time.Since(start))
		start1 := time.Now()
		result1 := numBusesToDestination1(test.routes, test.source, test.target)
		fmt.Println("自己写的", result1, test.expect, time.Since(start1))
	}
}
