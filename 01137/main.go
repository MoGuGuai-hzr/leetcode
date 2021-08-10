package main

// 最简单写法, 会超时
func tribonacci1(n int) int {
	switch n {
	case 0:
		return 0
	case 1, 2:
		return 1
	default:
		return tribonacci1(n-1) + tribonacci1(n-2) + tribonacci1(n-3)
	}
}

// 额外空间, 如果把 m 弄成全局或者放到某个结构体里面, 可以减少多次调用的时间
func tribonacci2(n int) int {
	m := make(map[int]int, 38)
	m[0] = 0
	m[1] = 1
	m[2] = 1

	maxI := 2
	if v, have := m[n]; have {
		return v
	} else {
		for i := maxI + 1; i <= n; i++ {
			m[i] = m[i-1] + m[i-2] + m[i-3]
		}
		return m[n]
	}
}

// 每次都计算, 节约了空间
func tribonacci3(n int) int {
	a, b, c := 0, 1, 1
	for i := 1; i <= n; i++ {
		c, b, a = c+b+a, c, b
	}
	return a
}

// 官方解答, 跪了
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	m := matrix{
		{1, 1, 1},
		{1, 0, 0},
		{0, 1, 0},
	}
	res := m.pow(n)
	return res[0][2]
}

type matrix [3][3]int

func (a matrix) mul(b matrix) matrix {
	c := matrix{}
	for i, row := range a {
		for j := range b[0] {
			for k, v := range row {
				c[i][j] += v * b[k][j]
			}
		}
	}
	return c
}

func (a matrix) pow(n int) matrix {
	res := matrix{}
	for i := range res {
		res[i][i] = 1
	}
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res.mul(a)
		}
		a = a.mul(a)
	}
	return res
}
