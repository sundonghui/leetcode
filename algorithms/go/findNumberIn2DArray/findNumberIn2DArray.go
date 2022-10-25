package findnumberin2darray

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, list := range matrix {
		r := binarySearch(list, target)
		if r == true {
			return true
		}
	}
	return false
}

func binarySearch(list []int, target int) bool {
	if len(list) < 1 {
		return false
	}
	if len(list) == 1 {
		return list[0] == target
	}
	mid := len(list) / 2
	if target == list[mid] {
		return true
	}
	if target > list[mid] {
		return binarySearch(list[mid:], target)
	}
	return binarySearch(list[:mid], target)
}
