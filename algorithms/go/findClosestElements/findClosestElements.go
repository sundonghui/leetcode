package findclosestelements

func findClosestElements(arr []int, k int, x int) []int {
	start := 0
	end := len(arr) - k

	for start < end {
		mid := start + (end-start)/2

		if x-arr[mid] > arr[mid+k]-x {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return arr[start : start+k]
}

// 寻找最接近 x 的 k 个值
func findClosestElements1(arr []int, k int, x int) []int {
	if x <= arr[0] {
		return arr[:k]
	}
	n := len(arr) - 1
	if x >= arr[n] {
		return arr[n-k+1:]
	}

	left, right := 0, n
	for left <= right {
		mid := (right + left) / 2
		if arr[mid] >= x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	pre, next := left-1, left
	for ; k > 0; k-- {
		if pre < 0 {
			next++
		} else if next > n || x-arr[pre] <= arr[next]-x {
			pre--
		} else {
			next++
		}
	}
	return arr[pre+1 : next]
}
