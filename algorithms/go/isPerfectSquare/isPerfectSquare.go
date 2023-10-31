package isperfectsquare

func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := (left + right) / 2
		sqrt := mid * mid
		if sqrt == num {
			return true
		}
		if sqrt > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
