package main

import "fmt"

func countPairs(deliciousness []int) int {
	m := make(map[int]int)
	ctn := 0
	list := make([]int, 22)
	for i := 0; i < 22; i++ {
		list[i] = 1 << i
	}
	for _, d := range deliciousness {
		for _, l := range list {
			if l >= d {
				ctn += m[l-d]
			}
		}
		m[d]++
	}
	return ctn % (1e9 + 7)
}

func main() {
	for _, test := range []struct {
		in  []int
		exp int
	}{
		{[]int{1, 3, 5, 7, 9}, 4},
		{[]int{1, 1, 1, 3, 3, 3, 7}, 15},
		{[]int{149, 107, 1, 63, 0, 1, 6867, 1325, 5611, 2581, 39, 89, 46, 18, 12, 20, 22, 234}, 12},
		{[]int{2160, 1936, 3, 29, 27, 5, 2503, 1593, 2, 0, 16, 0, 3860, 28908, 6, 2, 15, 49, 6246, 1946, 23, 105, 7996, 196, 0, 2, 55, 457, 5, 3, 924, 7268, 16, 48, 4, 0, 12, 116, 2628, 1468}, 53},
	} {
		rst := countPairs(test.in)
		fmt.Println(rst, rst == test.exp)
	}
}
