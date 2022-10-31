package numways

func numWays(n int) int {
	const mod int = 1e9 + 7
	p, q := 1, 1
	for i := 0; i < n; i++ {
		p, q = q, (p+q)%mod

	}
	return p
}
