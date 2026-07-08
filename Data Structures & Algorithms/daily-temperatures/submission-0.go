func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	
	for i := len(res)-2; i>=0; i-- {
		probe := i+1
		for {
			if temperatures[probe] > temperatures[i] {
				res[i] = probe - i
				break
			}
			if res[probe] == 0 {
				res[i] = 0
				break
			}
			probe = probe + res[probe]
		}
	}
	return res
}
