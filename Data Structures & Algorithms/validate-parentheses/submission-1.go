func isValid(s string) bool {
    sr := []rune(s)
	
	closer := map[rune]rune {
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []rune{}

	for _, r := range sr {
		if spec_closer, ok := closer[r]; ok {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == spec_closer {
				stack = stack[0:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, r)
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
