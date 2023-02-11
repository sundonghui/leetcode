package reverse

func reverse(x int) int {
	isPos := true
	if x < 0 {
		isPos = false
		x = x * -1
	}
	var res int
	for x >= 10 {
		n := x % 10
		res = res*10 + n
		if res > (1<<31 - 1) {
			return 0
		}
		x = x / 10
	}
	res = res*10 + x
	if res > (1<<31 - 1) {
		return 0
	}
	if isPos {
		return res
	}
	return -1 * res
}
