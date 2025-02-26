package rotate

func rotate(matrix [][]int) {
	n := len(matrix)
	// 水平旋转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 对角线旋转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// 旋转 90 度为什么等于水平旋转➕对角线旋转
func rotate90(matrix [][]int) {
	n := len(matrix)
	// 水平旋转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}
	// 对角线旋转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func Rotate(nums []int, k int) {
	list := make([]int, len(nums))
	for i, v := range nums {
		list[(i+k)%len(list)] = v
	}
	copy(nums, list)
}
