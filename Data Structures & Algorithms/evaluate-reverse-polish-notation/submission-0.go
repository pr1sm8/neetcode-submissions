func evalRPN(tokens []string) int {
	stack := []int{}
	for _, r := range tokens {
		if r == "+" {
			res := stack[len(stack) - 2] + stack[len(stack) - 1]
			stack = stack[0:len(stack)-2]
			stack = append(stack, res)
		} else if r == "-" {
			res := stack[len(stack) - 2] - stack[len(stack) - 1]
			stack = stack[0:len(stack)-2]
			stack = append(stack, res)
		} else if r == "*" {
			res := stack[len(stack) - 2] * stack[len(stack) - 1]
			stack = stack[0:len(stack)-2]
			stack = append(stack, res)
		} else if  r == "/" {
			res := stack[len(stack) - 2] / stack[len(stack) - 1]
			stack = stack[0:len(stack)-2]
			stack = append(stack, res)
		} else {
			val, _ := strconv.Atoi(r)
			stack = append(stack, val)
		}
	}
	return stack[0]
}
