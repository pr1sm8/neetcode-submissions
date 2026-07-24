func minEatingSpeed(piles []int, h int) int {
	lower := 1
	upper := 0
	for _, p := range piles {
		if p > upper {
			upper = p
		}
	}
	// for i := 1; i <= upper; i++ {
	// 	fmt.Println(i, solve(piles, i))
	// }
	// return 0
	for lower < upper {
		mid := (lower + upper ) / 2
		t := solve(piles, mid)
		if t > h {
			lower = mid + 1
		} else {
			upper = mid
		}
	}
	return lower	
}

func solve(piles []int, rate int) (time int) {
	for _, p := range piles {
		time += (p + rate - 1) / rate
	}
	return
}