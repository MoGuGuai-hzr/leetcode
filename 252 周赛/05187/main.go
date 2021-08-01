package main

import (
	"fmt"
	"sort"
)

func minimumPerimeter(neededApples int64) int64 {
	var sum int64
	var val int64
	var i int64
	for sum < neededApples {
		i++
		val += 24*i - 12
		sum += val
	}
	return 8 * i
}

// 官方解答: Sn​=2n(n+1)(2n+1)
func minimumPerimeter2(neededApples int64) int64 {
	var sn int64
	var n int64
	for sn < neededApples {
		n++
		sn = 2 * n * (n + 1) * (2*n + 1)
	}
	return n * 8
}

// 官方解答: Sn​=2n(n+1)(2n+1), 再利用二分查找
func minimumPerimeter3(neededApples int64) int64 {
	n := sort.Search(100000, func(i int) bool {
		return 2*int64(i)*(int64(i)+1)*(2*int64(i)+1) >= neededApples
	})

	return int64(n) * 8
}

func main() {
	for _, test := range []struct {
		neededApples int64
		exp          int64
	}{
		{1, 8},
		{13, 16},
		{60, 16},
		{1000000000, 5040},
		{100000000000000, 233920},
		{2784381467700, 70896},
	} {
		r := minimumPerimeter3(test.neededApples)
		fmt.Println(r, test.exp)
	}

}
