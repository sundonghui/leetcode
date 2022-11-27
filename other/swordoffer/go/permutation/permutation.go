package permutation

func permutation(s string) []string {
	res := []string{}
	bytes := []byte(s)
	var dfs func(start int)
	dfs = func(start int) {
		if start == len(bytes)-1 {
			res = append(res, string(bytes))
		}
		dict := map[byte]bool{}
		for i := start; i < len(bytes); i++ {
			if !dict[bytes[i]] {
				bytes[start], bytes[i] = bytes[i], bytes[start]
				dict[bytes[start]] = true
				dfs(start + 1)
				bytes[start], bytes[i] = bytes[i], bytes[start]
			}
		}
	}
	dfs(0)
	return res
}
