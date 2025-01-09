package groupanagrams

import "sort"

// func groupAnagrams(strs []string) [][]string {
// 	if len(strs) == 0 {
// 		return [][]string{}
// 	}
// 	if len(strs) == 1 {
// 		return [][]string{strs}
// 	}
// 	anagrams := make([]map[byte]int, 0)
// 	// map key is anagrams index
// 	m := make(map[int][]string)
// 	for _, str := range strs {
// 		strAga := getStrAnagram(str)
// 		if isAnagramExist(anagrams, strAga) {
// 			index := getAnagramIndex(anagrams, strAga)
// 			m[index] = append(m[index], str)
// 		} else {
// 			anagrams = append(anagrams, strAga)
// 			m[len(anagrams)-1] = []string{str}
// 		}
// 	}
// 	if len(m) == 0 {
// 		return [][]string{}
// 	}
// 	result := [][]string{}
// 	for _, v := range m {
// 		result = append(result, v)
// 	}
// 	return result
// }

// func getStrAnagram(str string) map[byte]int {
// 	m := make(map[byte]int)
// 	for i := 0; i < len(str); i++ {
// 		m[str[i]]++
// 	}
// 	return m
// }

// func isAnagramExist(anagrams []map[byte]int, anagram map[byte]int) bool {
// 	for _, a := range anagrams {
// 		if isSameMap(a, anagram) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func getAnagramIndex(anagrams []map[byte]int, anagram map[byte]int) int {
// 	for index, a := range anagrams {
// 		if isSameMap(a, anagram) {
// 			return index
// 		}
// 	}
// 	return -1
// }

// func isSameMap(anagram1, anagram2 map[byte]int) bool {
// 	if len(anagram1) != len(anagram2) {
// 		return false
// 	}
// 	for k, v := range anagram1 {
// 		if w, ok := anagram2[k]; !ok || w != v {
// 			return false
// 		}
// 	}
// 	return true
// }

// 先排序再处理
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		m[sortedStr] = append(m[sortedStr], str)
	}
	result := [][]string{}
	for _, value := range m {
		result = append(result, value)
	}
	return result
}
