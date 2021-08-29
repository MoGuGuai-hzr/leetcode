package main

import (
	"fmt"
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}

	for _, flight := range flights {
		graph[flight[0]][flight[1]] = flight[2]
	}

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[src] = 0

	queue := [][3]int{{src, 0, -1}}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for i, price := range graph[node[0]] {
			if price > 0 && node[1] <= dist[i] && node[2] < k && dist[i] > dist[node[0]]+price {
				dist[i] = node[1] + price
				queue = append(queue, [3]int{i, dist[i], node[2] + 1})
			}
		}
	}

	if dist[dst] == math.MaxInt32 {
		return -1
	}
	return dist[dst]
}

func main() {
	for _, test := range []struct {
		n, src, dst, k int
		flights        [][]int
		exp            int
	}{
		{3, 0, 2, 1, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 200},
		{3, 0, 2, 0, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 500},
		{3, 0, 2, -1, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, -1},
		{3, 1, 2, 1, [][]int{{0, 1, 2}, {1, 2, 1}, {2, 0, 10}}, 1},
		{4, 0, 3, 1, [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}, 6},
	} {
		r := findCheapestPrice(test.n, test.flights, test.src, test.dst, test.k)
		fmt.Println(r, r == test.exp)
	}
}
