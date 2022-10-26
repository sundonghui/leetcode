package findnumberin2darray

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, list := range matrix {
		isOk := binarySearch(list, target)
		if isOk {
			return isOk
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
	if list[0] > target || list[len(list)-1] < target {
		return false
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

func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		fmt.Println(x, y)
		v := matrix[x][y]
		fmt.Println(v)
		if v == target {
			return true
		} else if v > target {
			y--
		} else {
			x++
		}
	}
	return false
}
