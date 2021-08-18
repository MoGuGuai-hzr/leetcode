package main

func checkRecord(n int) (ans int) {
	const mod int = 1e9 + 7
	dp := [2][3]int{} // A 的数量，结尾连续 L 的数量
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		dpNew := [2][3]int{}
		// 以 P 结尾的数量
		for j := 0; j <= 1; j++ {
			for k := 0; k <= 2; k++ {
				dpNew[j][0] = (dpNew[j][0] + dp[j][k]) % mod
			}
		}
		// 以 A 结尾的数量
		for k := 0; k <= 2; k++ {
			dpNew[1][0] = (dpNew[1][0] + dp[0][k]) % mod
		}
		// 以 L 结尾的数量
		for j := 0; j <= 1; j++ {
			for k := 1; k <= 2; k++ {
				dpNew[j][k] = (dpNew[j][k] + dp[j][k-1]) % mod
			}
		}
		dp = dpNew
	}
	for j := 0; j <= 1; j++ {
		for k := 0; k <= 2; k++ {
			ans = (ans + dp[j][k]) % mod
		}
	}
	return ans
}
