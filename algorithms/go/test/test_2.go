package test

import "fmt"

// Whale 帷幄
// 找出字符串内有效且连续的大括号的最长子串

// 题目描述
// 要求 大括号输出有效且连续​
// 输入 "{}{{}{}}{{{}}}{{}{{}}}"​

// 输出是 "{{{}}}"​

// 可以认为输入是可信的。长度大于0 且仅有 '{' 和 '}'

func main2(input string) string {
	//output := "{{{}}}"
	size, max := 0, 0
	maxStr := ""
	for i, v := range input {
		if v == '{' {
			size++
			if size > max {
				max = size
				maxStr = getStr(input, i-max, i+max)
			}
		} else {
			size--
		}
	}
	return maxStr
}

func getStr(str string, i, j int) string {
	res := ""
	for k, v := range str {
		if k > i && k <= j {
			res = fmt.Sprintf("%s%s", res, string(v))
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxStrFunc(i, j string) string {
	if len(i) > len(j) {
		return i
	}
	return j
}
