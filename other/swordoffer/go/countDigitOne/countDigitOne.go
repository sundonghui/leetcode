package countdigitone

func countDigitOne(n int) int {
	ans := 0
	for decimal := 1; n >= decimal; {
		ans += (n/(decimal*10))*decimal + min(max(n%(decimal*10)-decimal+1, 0), decimal)
		decimal *= 10
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
