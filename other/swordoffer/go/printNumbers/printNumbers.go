package printnumbers

func printNumbers(n int) []int {
	max := quickCal(10, n)
	list := make([]int, 0, max)
	for i := 1; i < max; i++ {
		list = append(list, i)
	}
	return list
}

func quickCal(n, y int) int {
	if y == 0 {
		return 1
	}
	x := quickCal(n, y/2)
	if y%2 == 0 {
		return x * x
	}
	return n * x * x
}
