package main

import "fmt"

func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs) + 1
	m := make(map[int][]int, n)
	rst := make([]int, n)
	for _, v := range adjacentPairs {
		if _, ok := m[v[0]]; !ok {
			m[v[0]] = make([]int, 0, 2)
		}
		m[v[0]] = append(m[v[0]], v[1])
		if _, ok := m[v[1]]; !ok {
			m[v[1]] = make([]int, 0, 2)
		}
		m[v[1]] = append(m[v[1]], v[0])
	}

	for k, v := range m {
		if len(v) == 1 {
			rst[0] = k
			rst[1] = v[0]
			break
		}
	}

	for i, v := 2, m[rst[1]]; i < n; i++ {
		if rst[i-2] != v[0] {
			rst[i] = v[0]
			v = m[v[0]]
		} else {
			rst[i] = v[1]
			v = m[v[1]]
		}
	}

	return rst
}

func main() {
	for _, test := range []struct {
		adjacentPairs [][]int
		exp           []int
	}{
		{[][]int{{2, 1}, {3, 4}, {3, 2}}, []int{1, 2, 3, 4}},
	} {
		r := restoreArray(test.adjacentPairs)
		fmt.Println(r, test.exp)
	}
}
