func minWindow(s string, t string) string {
	sr := []rune(s)
	tr := []rune(t)
	trh := make(map[rune]int)
	for _, v := range tr {
		trh[v]++
	}

	windowh := make(map[rune]int)
	result := ""

	p0 := 0
	p1 := -1
	for {
		// fmt.Println(p0, p1, isvalid(trh, windowh))

		valid := true
		for k, v := range trh {
			if windowh[k] < v {
				valid = false
				break
			}
		}

		if valid {
			if p1-p0+1 < len(result) || result == "" {
				result = string(sr[p0:p1+1])
			}
			windowh[ sr[p0] ]--
			p0++
		} else {
			if p1 == len(sr)-1 {
				break
			}
			p1++
			windowh[ sr[p1] ]++
		}
	}
	return result
}

// func isvalid(trh, windowh map[rune]int) bool {
// 	for k, v := range trh {
// 		if windowh[k] < v {
// 			return false
// 		}
// 	}
// 	return true
// }
