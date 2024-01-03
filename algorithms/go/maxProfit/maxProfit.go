package maxprofit

func maxProfit(prices []int) int {
	minValue := prices[0]
	maxValue := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minValue {
			minValue = prices[i]
		}
		if prices[i]-minValue > maxValue {
			maxValue = prices[i] - minValue
		}
	}
	return maxValue
}
