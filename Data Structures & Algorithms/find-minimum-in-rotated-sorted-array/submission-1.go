func findMin(nums []int) int {
	lower := 0 
	upper := len(nums) - 1
	for upper - lower > 1 {
		mid := (lower+upper) / 2
		if nums[mid] < nums[upper] {
			upper = mid
		} else {
			lower = mid
		}
	}
	ans := nums[lower]
	if nums[upper] < ans {
		ans = nums[upper]
	}
	return ans
}
