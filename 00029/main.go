package main

import "fmt"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	sign := true
	if dividend < 0 {
		sign = !sign
		dividend = -dividend
	}
	if divisor < 0 {
		sign = !sign
		divisor = -divisor
	}

	nums := divisor
	i := 0
	for i = 0; dividend-nums >= 0; i++ {
		nums = nums << 1
	}
	cnt := 0
	for i > 0 {
		nums = nums >> 1
		i--
		if dividend-nums >= 0 {
			dividend -= nums
			cnt += 1 << i
		}
	}

	if sign {
		if cnt > 1<<31-1 {
			return 1<<31 - 1
		}
		return cnt
	}
	if cnt > 1<<31 {
		return 1<<31 - 1
	}
	return -cnt
}

func main() {
	for _, test := range []struct {
		dividend, divisor, exp int
	}{
		{10, 3, 3},
		{15, 3, 5},
		{-2147483648, -1, 2147483647},
		{-2147483648, 1, -2147483648},
		{-2147483649, 1, 2147483647},
	} {
		r := divide(test.dividend, test.divisor)
		fmt.Println(r, r == test.exp)
	}
}
