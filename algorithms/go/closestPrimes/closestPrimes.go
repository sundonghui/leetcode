package closestprimes

import "fmt"

func closestPrimes(left int, right int) []int {
	list := make([]int, 0, 2)
	result := make([]int, 2)
	min := 0
	for left <= right {
		if isPrime(left) {
			if len(list) <= 1 {
				list = append(list, left)
			}
			if len(list) == 2 {
				fmt.Println(list)
				if list[1] != left {
					list[0] = list[1]
					list[1] = left
				}
				tmp := list[1] - list[0]
				if tmp <= 2 {
					return []int{list[0], list[1]}
				}
				if tmp < min || min == 0 {
					min = tmp
					result[0] = list[0]
					result[1] = list[1]
				}
			}
		}
		left++
	}
	if result[0] == 0 {
		return []int{-1, -1}
	}
	return result
}

func closestPrimes1(left int, right int) []int {
	list := make([]int, 0, right-left)
	for left <= right {
		if isPrime(left) {
			list = append(list, left)
		}
		left++
	}
	fmt.Println(list)
	if len(list) < 2 {
		return []int{-1, -1}
	}
	pre, cur := 0, 1
	min := list[1] - list[0]
	for i := 1; i < len(list); i++ {
		pre, cur = i-1, i
		tmp := list[cur] - list[pre]
		// 如果等于 2，一定就是答案
		if tmp <= 2 {
			return []int{list[pre], list[cur]}
		}
		if tmp < min {
			min = tmp
		}
	}
	return []int{list[pre], list[cur]}
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
