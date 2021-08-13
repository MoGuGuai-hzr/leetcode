package main

import "fmt"

// 使用迭代代替递归
func countDigitOne(n int) int {
	base := []int{0, 1, 20, 300, 4000, 50000, 600000, 7000000, 80000000, 900000000, 1000000000}
	cnt := 0

	number := fmt.Sprintf("%d", n)

	var f func([]byte)
	f = func(b []byte) {
		nn := len(b)
		if nn == 1 {
			if b[0] >= '1' {
				cnt++
			}
			return
		}

		switch b[0] {
		case '0':
		case '1':
			cnt += base[nn-1]
			for i, t := nn-1, 1; i > 0; i-- {
				cnt += int(b[i]-'0') * t
				t *= 10
			}
			cnt += 1
		default:
			cnt += base[nn-1] * int(b[0]-'1')
			cnt += base[nn] - base[nn-1]*9
		}
	}

	for i := 0; i < len(number); i++ {
		f([]byte(number)[i:])
	}

	fmt.Println(cnt)
	return cnt
}

// 初始版本
func countDigitOne1(n int) int {
	base := []int{0, 1, 20, 300, 4000, 50000, 600000, 7000000, 80000000, 900000000, 1000000000}
	cnt := 0

	number := fmt.Sprintf("%d", n)

	var f func([]byte)
	f = func(b []byte) {
		nn := len(b)
		if nn == 1 {
			if b[0] >= '1' {
				cnt++
			}
			return
		}

		switch b[0] {
		case '0':
		case '1':
			cnt += base[nn-1]
			for i, t := nn-1, 1; i > 0; i-- {
				cnt += int(b[i]-'0') * t
				t *= 10
			}
			cnt += 1
		default:
			cnt += base[nn-1] * int(b[0]-'1')
			cnt += base[nn] - base[nn-1]*9
		}
		f(b[1:])
	}

	f([]byte(number))

	fmt.Println(cnt)
	return cnt
}

// 官方解答
func countDigitOne0(n int) (ans int) {
	// mulk 表示 10^k
	// 在下面的代码中，可以发现 k 并没有被直接使用到（都是使用 10^k）
	// 但为了让代码看起来更加直观，这里保留了 k
	for k, mulk := 0, 1; n >= mulk; k++ {
		ans += (n/(mulk*10))*mulk + min(max(n%(mulk*10)-mulk+1, 0), mulk)
		mulk *= 10
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	countDigitOne(150)
	countDigitOne(101)
	countDigitOne(13)
}
