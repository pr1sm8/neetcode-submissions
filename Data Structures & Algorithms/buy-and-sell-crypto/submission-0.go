func maxProfit(prices []int) int {
	// minp := make([]int, len(prices))
	maxp := make([]int, len(prices))

	// minp[0] = prices[0]
	ln := len(prices)-1
	maxp[ln] = prices[ln]

	for i := 1; i < len(prices); i++ {
		// minp[i] = min(minp[i-1], prices[i])
		maxp[ln-i] = max(maxp[ln-i+1], prices[ln-i])
	}
	// fmt.Println(maxp)
	maxRes := 0
	for i := range prices {
		maxRes = max(maxRes, maxp[i] - prices[i])
	}
	return maxRes
}
