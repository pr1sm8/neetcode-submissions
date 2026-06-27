type MinStack struct {
	valstack []int
	prefixminstack []int
}

func Constructor() MinStack {
	return MinStack{
		valstack: []int{},
		prefixminstack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.valstack = append(this.valstack, val)
	if len(this.prefixminstack) == 0 {
		this.prefixminstack = append(this.prefixminstack, val)
	} else {
		beforecur := this.prefixminstack[len(this.prefixminstack) - 1]
		if val < beforecur {
			beforecur = val
		}
		this.prefixminstack = append(this.prefixminstack, beforecur)
	}
}

func (this *MinStack) Pop() {
	this.valstack = this.valstack[0:len(this.valstack) - 1]
	this.prefixminstack = this.prefixminstack[0:len(this.prefixminstack) - 1]
}

func (this *MinStack) Top() int {
	return this.valstack[len(this.valstack) - 1]
}

func (this *MinStack) GetMin() int {
	return this.prefixminstack[len(this.prefixminstack) - 1]
}
