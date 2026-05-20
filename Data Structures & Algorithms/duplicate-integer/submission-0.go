func hasDuplicate(nums []int) bool {
    dupMap := make(map[int]struct{})
    for _, val := range nums {
        _, ok := dupMap[val]
        if ok {
            return true
        }
        dupMap[val] = struct{}{}

    }
    return false
}
