package main

import "fmt"

func allPathsSourceTarget(graph [][]int) (ans [][]int) {
	stk := []int{0}
	var dfs func(int)
	dfs = func(x int) {
		if x == len(graph)-1 {
			ans = append(ans, append([]int(nil), stk...))
			return
		}
		for _, y := range graph[x] {
			stk = append(stk, y)
			dfs(y)
			stk = stk[:len(stk)-1]
		}
	}
	dfs(0)
	return
}

func main() {
	for _, test := range []struct {
		graph [][]int
		exp   [][]int
	}{
		{[][]int{{1, 2}, {3}, {3}, {}}, [][]int{{0, 1, 3}, {0, 2, 3}}},
	} {
		r := allPathsSourceTarget(test.graph)
		fmt.Println(r, compare(r, test.exp))
	}
}

func compare(x, y [][]int) bool {
	if len(x) != len(y) {
		return false
	}

	for i := 0; i < len(x); i++ {
		if len(x[i]) != len(y[i]) {
			return false
		}
		for j := 0; j < len(x[i]); j++ {
			if x[i][j] != y[i][j] {
				return false
			}
		}
	}

	return true
}
