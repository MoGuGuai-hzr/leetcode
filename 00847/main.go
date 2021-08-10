package main

import (
	"math"
	"math/bits"
)

func shortestPathLength1(graph [][]int) int {
	n := len(graph)
	type tuple struct{ u, mask, dist int }
	q := []tuple{}
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, 1<<n)
		seen[i][1<<i] = true
		q = append(q, tuple{i, 1 << i, 0})
	}

	for {
		t := q[0]
		q = q[1:]
		if t.mask == 1<<n-1 {
			return t.dist
		}
		// 搜索相邻的节点
		for _, v := range graph[t.u] {
			maskV := t.mask | 1<<v
			if !seen[v][maskV] {
				q = append(q, tuple{v, maskV, t.dist + 1})
				seen[v][maskV] = true
			}
		}
	}
}

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			d[i][j] = n + 1
		}
	}
	for v, nodes := range graph {
		for _, u := range nodes {
			d[v][u] = 1
		}
	}

	// 使用 floyd 算法预处理出所有点对之间的最短路径长度
	for k := range d {
		for i := range d {
			for j := range d {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, 1<<n)
		for j := range f[i] {
			f[i][j] = math.MaxInt64 / 2
		}
	}
	for mask := 1; mask < 1<<n; mask++ {
		// 如果 mask 只包含一个 1，即 mask 是 2 的幂
		if mask&(mask-1) == 0 {
			i := bits.TrailingZeros(uint(mask))
			f[i][1<<i] = 0
			continue
		}
		for u := 0; u < n; u++ {
			if mask>>u&1 > 0 {
				for v := 0; v < n; v++ {
					if v != u && mask>>v&1 > 0 {
						f[u][mask] = min(f[u][mask], f[v][mask^(1<<u)]+d[v][u])
					}
				}
			}
		}
	}
	ans := math.MaxInt64
	for u := 0; u < n; u++ {
		ans = min(ans, f[u][1<<n-1])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
