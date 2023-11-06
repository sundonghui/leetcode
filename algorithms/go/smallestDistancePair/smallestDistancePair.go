package smallestdistancepair

import "sort"

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	//第 k 小的距离取值范围
	left, right := 0, nums[len(nums)-1]-nums[0]
	for left < right {
		//二分查找
		mid := left + (right-left)>>1
		var count int
		//统计满足距离小于等于 mid 的数对个数 count
		for i, j := 0, 0; j < len(nums); j++ {
			//不满足条件时，收缩左边界
			for nums[j]-nums[i] > mid {
				i++
			}
			count += j - i
		}
		//二分查找左侧边界，防止找到不存在的数
		if count >= k {
			right = mid
		}
		if count < k {
			left = mid + 1
		}
	}
	return right
}

func mock(list []int, k int) int {
	sort.Ints(list)
	left, right := 0, list[len(list)-1]-list[0]
	for left < right {
		mid := left + (right-left)>>1
		var count int
		for i, j := 0, 0; j < len(list); j++ {
			for list[j]-list[i] > mid {
				i++
			}
			count += j - i
		}
		if count >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
