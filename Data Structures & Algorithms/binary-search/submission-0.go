func search(nums []int, target int) int {
	return searchb(nums, target, 0, len(nums)-1)
}

func searchb(nums[]int, target, start, end int) int {
	if start > end {
		return -1
	}
	if start == end {
		if nums[start] == target {
			return start
		} else {
			return -1
		}
	}
	mid := (start+end) /2
	if nums[mid] == target {
		return mid
	} else if target < nums[mid]{
		return searchb(nums, target, start, mid-1)
	} else {
		return searchb(nums, target, mid+1, end)
	}
}
