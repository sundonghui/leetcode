package isisomorphic

func isIsomorphic(s string, t string) bool {
	m1 := make(map[byte]byte, len(s))
	m2 := make(map[byte]byte, len(s))
	tb1 := make([]byte, 0, len(s))
	tb2 := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		b1 := s[i]
		if v, ok := m1[b1]; !ok {
			m1[b1] = t[i]
		} else if v != m1[b1] {
			return false
		}
		b2 := t[i]
		if v, ok := m2[b2]; !ok {
			m2[b2] = s[i]
		} else if v != m2[b2] {
			return false
		}
		tb1 = append(tb1, m1[b1])
		tb2 = append(tb2, m2[b2])
	}
	return string(tb1) == t && string(tb2) == s
}

func isIsomorphic_official(s, t string) bool {
	s2t := map[byte]byte{}
	t2s := map[byte]byte{}
	for i := range s {
		x, y := s[i], t[i]
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}
