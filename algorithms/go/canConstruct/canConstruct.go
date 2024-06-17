package canconstruct

func canConstruct(ransomNote string, magazine string) bool {
	mm := make(map[byte]int, len(magazine))
	for i := 0; i < len(magazine); i++ {
		mm[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		if cnt, ok := mm[ransomNote[i]]; ok && cnt > 0 {
			mm[ransomNote[i]]--
		} else {
			return false
		}
	}
	return true
}
