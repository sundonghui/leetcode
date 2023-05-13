package generateparenthesis

var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0, n)
	dfs(n, 0, 0, "")
	return res
}

func dfs(n, lc, rc int, path string) {
	if lc == n && rc == n {
		res = append(res, path)
		return
	}
	if lc < n {
		dfs(n, lc+1, rc, path+"(")
	}
	if rc < lc {
		dfs(n, lc, rc+1, path+")")
	}
}
