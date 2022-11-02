package cuttingrope

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = max(curMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func cuttingRope2(n int) int {
	mod := int(1e9 + 7)
	if n <= 3 {
		return n - 1
	}
	r := 1
	for n > 4 {
		r = r * 3 % mod
		n -= 3
	}
	return r * n % mod
}
