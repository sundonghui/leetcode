package twosum

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for index, v := range nums {
		m[target-v] = index
	}

	for index, v := range nums {
		if v, ok := m[v]; ok {
			if v != index {
				return []int{v, index}
			}
		}
	}
	return nil
}

func twoSum(numbers []int, target int) []int {
	for i := 0; i <= len(numbers)-1; i++ {
		left, right := i+1, len(numbers)-1
		for left <= right {
			mid := left + (right-left)/2
			if numbers[i]+numbers[mid] == target {
				return []int{i + 1, mid + 1}
			}
			if numbers[mid] > target-numbers[i] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return []int{0, 0}
}

func twoSumTwoPointer(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		left, right := i+1, len(numbers)-1
		for left <= right {
			mid := left + (right-left)/2
			if numbers[i]+numbers[mid] == target {
				return []int{i + 1, mid + 1}
			}
			if numbers[mid]+numbers[i] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return []int{}
}
