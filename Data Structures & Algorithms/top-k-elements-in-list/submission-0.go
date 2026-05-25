func topKFrequent(nums []int, k int) []int {
	counts := map[int]int{}
	for _, n := range nums {
		counts[n]++
	}
	freq := make([][]int, len(nums)+1)
	for k, v := range counts {
		freq[v] = append(freq[v], k)
	}
	res := []int{}
	for i := len(freq)-1; i>=0; i -- {
		for _, val := range freq[i] {
			res = append(res, val)
			if len(res) == k {
				return res
			}
		}
	}
	return res
	
}
