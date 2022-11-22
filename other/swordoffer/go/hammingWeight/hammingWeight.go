package hammingweight

func hammingWeight(num uint32) int {
	if num < 1 {
		return 0
	}

	res := 0
	for num > 0 {
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}

func hammingWeight1(num uint32) int {
	if num < 1 {
		return 0
	}

	res := 0
	for num > 0 {
		if num%2 == 1 {
			res++
		}
		num /= 2
	}
	return res
}
