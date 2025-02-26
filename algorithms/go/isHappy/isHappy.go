package ishappy

func isHappy(n int) bool {
	slow, fast := n, step(n)
	for slow != fast && fast != 1 {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
