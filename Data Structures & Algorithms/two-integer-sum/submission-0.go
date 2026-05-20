func twoSum(nums []int, target int) []int {
    indexMap := make(map[int]int)
	for i, num := range nums {
		comp := target-num
		compIndex, ok := indexMap[comp]
		if ok {
			return []int{compIndex, i}
		} else {
			_, ok_num := indexMap[num]
			if !ok_num {
				indexMap[num] = i
			}
		}
	}
	return []int{}
}
