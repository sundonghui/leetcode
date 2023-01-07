package findcontinuoussequence

func findContinuousSequence(target int) [][]int {
	i, j, sum := 1, 1, 0
	ret := make([][]int, 0)
	for i <= target>>1 {
		if sum < target {
			sum += j
			j++
		} else if sum > target {
			sum -= i
			i++
		} else {
			tmp := make([]int, 0, j-1)
			for m := i; m < j; m++ {
				tmp = append(tmp, m)
			}
			ret = append(ret, tmp)
			sum -= i
			i++
		}
	}
	return ret
}
