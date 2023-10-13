package findthearrayconcval

import (
	"fmt"
	"strconv"
)

func findTheArrayConcVal(nums []int) int64 {
	start, end := 0, len(nums)-1
	var result int64
	for start <= end {
		if start == end {
			result += int64(nums[start])
			break
		}
		tmp := fmt.Sprintf("%d%d", nums[start], nums[end])
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return 0
		}
		result = result + tmpInt64
		start++
		end--
	}
	return result
}
