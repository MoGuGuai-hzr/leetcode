package main

import (
	"fmt"
)

type passed struct {
	bus, road []int
}

func InSlice(list []int, x int) bool {
	for _, v := range list {
		if v == x {
			return true
		}
	}
	return false
}

func numBusesToDestination(routes [][]int, source int, target int) int {
	ps := make([]passed, 0)
	for bus, list := range routes {
		if InSlice(list, source) {
			if InSlice(list, target) {
				return 1
			} else {
				p := passed{bus: []int{bus}, road: list}
				ps = append(ps, p)
			}
		}
	}
	for i := 0; i < len(routes); i++ {
	}
	return -1
}

func main() {
	for _, test := range []struct {
		routes                 [][]int
		source, target, expect int
	}{
		{[][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6, 2},
		{[][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12, -1},
	} {
		result := numBusesToDestination(test.routes, test.source, test.target)
		fmt.Println(result, test.expect)
	}
}
