func isAnagram(s string, t string) bool {
	totalCount := 0
	count := make(map[rune]int)
	sRune := []rune(s)
	for _, sr := range sRune {
		c, _ := count[sr]
		c++
		count[sr] = c
	}
	totalCount = len(sRune)
	sRune = []rune(t)
	for _, tr := range sRune {
		c, ok := count[tr]
		if !ok {
			return false
		}
		if c == 0 {
			return false
		}
		c--
		count[tr] = c
		totalCount--
	}
	if totalCount != 0 {
		return false
	}
	return true

}
