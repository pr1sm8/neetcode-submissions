func trap(height []int) int {
	leftHill := make([]int, len(height))
	rightHill := make([]int, len(height))
	leftHill[0] = height[0]
	rightHill[len(height)-1] = height[len(height)-1]
	for i:=1; i<len(height); i++ {
		leftHill[i] = max(leftHill[i-1], height[i])
		rightHill[len(height)-1-i] = max(rightHill[len(height)-i], height[len(height)-1 - i])
	}
	trapWater := 0
	// fmt.Println(leftHill)
	// fmt.Println(rightHill)
	for i := range height {
		trapWater += max(0, min(leftHill[i], rightHill[i]) - height[i])
	}
	return trapWater
}
