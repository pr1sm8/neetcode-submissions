func twoSum(numbers []int, target int) []int {
    p0 := 0
    p1 := len(numbers) - 1
    for p0 < p1 {
        sum := numbers[p0] + numbers[p1]
        if sum == target {
            return []int{p0+1, p1+1}
        } else if sum < target {
            p0 ++
        } else {
            p1 --
        }
    }
    return []int{0, 0}
}
