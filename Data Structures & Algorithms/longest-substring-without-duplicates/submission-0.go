func lengthOfLongestSubstring(s string) int {
	p0 := 0
	p1 := 0
	pres := make(map[rune]struct{})
	sr := []rune(s)
	maxlen := 0
	for p1 < len(sr) {
		if _, ok := pres[sr[p1]]; ok {
			for {
				if sr[p0] != sr[p1] {
					delete(pres, sr[p0])
					p0++
				} else {
					delete(pres, sr[p0])
					p0++
					break
				}
			}
		} else {
			pres[sr[p1]] = struct{}{}
			p1++
		}

		maxlen = max(maxlen, p1-p0)
	}
	return maxlen
}
