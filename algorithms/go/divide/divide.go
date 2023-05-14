package divide

import "math"

func divideTooStupid(dividend int, divisor int) int {
	flag := true
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		flag = false
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	max := 1<<31 - 1
	r := 0
	for dividend >= divisor {
		dividend -= divisor
		r++
		if r >= max {
			if flag {
				return r
			}
			return -r
		}
	}
	if flag {
		return r
	}
	return -r
}

func divide(dividend int, divisor int) int {
	res := 0
	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}

	dividend = abs(dividend)
	divisor = abs(divisor)

	for dividend >= divisor {
		d, q := divisor, 1
		for d <= dividend {
			d, q = d<<1, q<<1
		}
		res += q >> 1
		dividend -= (d >> 1)
		if res > math.MaxInt32 {
			if sign > 0 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
	}
	return sign * res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
