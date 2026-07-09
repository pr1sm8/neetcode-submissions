import "slices"
func carFleet(target int, position []int, speed []int) int {
	indices := make([]int, len(position))
	for i := range position {
		indices[i] = i
	}
	slices.SortFunc(indices, func(i, j int) int {
		return position[i] - position[j]
	})
	// fmt.Println(indices)
	// return 0
	

	newpos := make([]int, len(position))
	newspeed := make([]int, len(position))
	time := make([]float32, len(position))
	for i := range indices {
		newpos[i] = position[indices[i]]
		newspeed[i] = speed[indices[i]]
		time[i] = float32(target - newpos[i]) / float32(newspeed[i])
	}
	// fmt.Println(newpos)
	// fmt.Println(newspeed)

	toptime := time[len(position) - 1]
	ans := 1
	for i := len(position)-2; i>=0; i-- {
		if time[i] > toptime {
			toptime = time[i]
			ans++
		} 
	}

	return ans
}
