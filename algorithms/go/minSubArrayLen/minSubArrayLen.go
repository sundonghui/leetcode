package minsubarraylen

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	start, end := 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= target {
			ans = min(end-start+1, ans)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minSubArrayLen_copy(target int, nums []int) int {
	l, sum, ans := 0, 0, len(nums)+1 // 初始化 ans 应该 数组长度+1
	for r, x := range nums {         // 滑动数组 r 为结尾
		sum += x            // 滑动数组向右扩展
		for sum >= target { // 题中的大于等于
			ans = min(ans, r-l+1)
			sum -= nums[l]
			l++
		}
	}
	if ans > len(nums) { // 如果ans没有变化证明数组中所有值加到一起 < target
		return 0
	}
	return ans
}
