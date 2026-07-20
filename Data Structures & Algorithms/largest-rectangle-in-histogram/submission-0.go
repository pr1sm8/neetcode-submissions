func largestRectangleArea(heights []int) int {
	type pair struct {
		val int
		ind int
	}

	// calculating left to right for right projection
	rightproj := make([]int, len(heights))
	stack := make([]pair, 0)
	// stack = append(stack, pair{val: heights[0], ind: 0})

	for i := 0; i<len(heights); i++ {
		for len(stack) > 0 && heights[i] < stack[len(stack)-1].val {
			toppair := stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			rightproj[toppair.ind] = i - toppair.ind - 1
		}
		stack = append(stack, pair{val: heights[i], ind: i})
	}

	for i := range stack {
		rightproj[stack[i].ind] = len(heights) - stack[i].ind - 1
	}

	// fmt.Println("rightproj", rightproj)
	// calculating right to left for left projection
	leftproj := make([]int, len(heights))
	stack = make([]pair, 0)
	for i:=len(heights)-1; i>=0; i-- {
		for len(stack) > 0 && heights[i] < stack[len(stack)-1].val {
			toppair := stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			leftproj[toppair.ind] = toppair.ind - i - 1
		}
		stack = append(stack, pair{val: heights[i], ind: i})
	}
	for i := range stack {
		leftproj[stack[i].ind] = stack[i].ind
	}
	// fmt.Println("leftproj", leftproj)
	max := 0
	for i := range heights {
		calc := heights[i] * (rightproj[i] + leftproj[i] + 1)
		if calc > max {
			max = calc
		}
	}
	return max
}
