package main

func eventualSafeNodes(graph [][]int) (ans []int) {
	n := len(graph)
	color := make([]int, n)
	var safe func(int) bool
	safe = func(x int) bool {
		if color[x] > 0 {
			return color[x] == 2
		}
		color[x] = 1
		for _, y := range graph[x] {
			if !safe(y) {
				return false
			}
		}
		color[x] = 2
		return true
	}
	for i := 0; i < n; i++ {
		if safe(i) {
			ans = append(ans, i)
		}
	}
	return
}

func eventualSafeNodes3(graph [][]int) (ans []int) {
	n := len(graph)
	rg := make([][]int, n)
	inDeg := make([]int, n)
	for x, ys := range graph {
		for _, y := range ys {
			rg[y] = append(rg[y], x)
		}
		inDeg[x] = len(ys)
	}

	q := []int{}
	for i, d := range inDeg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		y := q[0]
		q = q[1:]
		for _, x := range rg[y] {
			inDeg[x]--
			if inDeg[x] == 0 {
				q = append(q, x)
			}
		}
	}

	for i, d := range inDeg {
		if d == 0 {
			ans = append(ans, i)
		}
	}
	return
}
