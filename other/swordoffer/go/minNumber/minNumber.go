package minnumber

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// minNumber 测试用例 23，输入 [20,1]，输出 102，不符合预期 120
func minNumber(nums []int) string {
	var sb strings.Builder
	for _, v := range nums {
		sb.WriteString(strconv.Itoa(v))
	}
	str := sb.String()
	list := make([]int, 0, len(str))
	for i := 0; i < len(str); i++ {
		num, _ := strconv.Atoi(string(str[i]))
		list = append(list, num)
	}
	sort.Ints(list)
	if len(list) > 0 {
		if len(list) == 1 {
			return strconv.Itoa(list[0])
		}
		first := 0
		for i, v := range list {
			if v > 0 {
				first = v
				list = append(list[:i], list[i+1:]...)
				break
			}
		}
		if first > 0 {
			var r strings.Builder
			r.WriteString(strconv.Itoa(first))
			for _, v := range list {
				r.WriteString(strconv.Itoa(v))
			}
			return r.String()
		}
	}
	return ""
}

// MinNumber 通过版
func MinNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x := fmt.Sprintf("%d%d", nums[i], nums[j])
		y := fmt.Sprintf("%d%d", nums[j], nums[i])
		return x < y
	})
	var ans string
	for _, n := range nums {
		ans += fmt.Sprintf("%d", n)
	}
	return ans
}
