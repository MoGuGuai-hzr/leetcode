package main

import "fmt"

func numWays(n int, relation [][]int, k int) int {
	ans := 0
	edges := make([][]int, n)
	for _, r := range relation {
		src, dst := r[0], r[1]
		edges[src] = append(edges[src], dst)
	}
	var dfs func(int, int)
	dfs = func(x, step int) {
		if step == k-1 {
			for _, y := range edges[x] {
				if y == n-1 {
					ans++
				}
			}
			return
		}
		for _, y := range edges[x] {
			dfs(y, step+1)
		}
	}
	dfs(0, 0)
	return ans
}

func main() {
	for _, test := range []struct {
		relation     [][]int
		n, k, expect int
	}{
		{[][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}}, 5, 3, 3},
		{[][]int{{0, 2}, {2, 1}}, 3, 2, 0},
	} {
		fmt.Println(numWays(test.n, test.relation, test.k), test.expect)
	}
}
