func characterReplacement(s string, k int) int {
	sr := []rune(s)
	p0 := 0 
	p1 := 0
	set := make(map[rune]int)
	maxlen := 0
	for p1 < len(sr) {
		if isValid(set, k) {
			maxlen = max(maxlen, p1-p0)
			set[sr[p1]]++
			p1++
		} else {
			set[sr[p0]]--
			if set[sr[p0]] == 0 {
				delete(set, sr[p0])
			}
			p0++
		}
	}
	if isValid(set, k) {
		maxlen = max(maxlen, p1-p0)
	}
	return maxlen
}

func isValid(set map[rune]int, k int) bool {
	total := 0
	max := 0
	for _, v := range set {
		total += v
		if max < v {
			max = v
		}
	}
	if total - max <= k {
		// fmt.Println(set, true)
		return true
	}
	// fmt.Println(set, false)
	return false
}
