package maxprofit

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n)
	for i := range dp {
		dp[i] = [2]int{}
	}
	dp[0][0] = 0
	dp[0][1] = 0 - prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
