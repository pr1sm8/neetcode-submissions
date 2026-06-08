func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	s1set := make(map[rune]int)
	s1r := []rune(s1)
	for _, v := range s1r {
		s1set[v]++
	}
	s2set := make(map[rune]int)
	s2r := []rune(s2)
	
	for i := range len(s1r) {
		s2set[s2r[i]]++
	}

	p2 := len(s1r)
	
	for p2 < len(s2r) {
		if isValid(s1set, s2set) {
			// fmt.Println(p2, true, s1set, s2set)
			return true
		}
		s2set[s2r[p2-len(s1r)]]--
		if s2set[s2r[p2-len(s1r)]] == 0 {
			delete(s2set, s2r[p2 - len(s1r)])
		}
		s2set[s2r[p2]]++
		p2++
	}
	if isValid(s1set, s2set) {
		// fmt.Println(p2, true, s1set, s2set)
		return true
	}
	return false
}

func isValid(s1set, s2set map[rune]int) bool {
	if len(s1set) != len(s2set) {
		return false
	}
	for k, v := range s1set {
		if s2set[k] != v {
			return false
		}
	}
	return true
}
