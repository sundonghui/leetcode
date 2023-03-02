package lettercombinations

var m = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var res []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	resLen := 1
	for i := 0; i < len(digits); i++ {
		resLen *= 4
	}
	res = make([]string, 0, resLen)
	backtrack(digits, 0, "")
	return res
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		res = append(res, combination)
	} else {
		digit := string(digits[index])
		letters := m[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}
