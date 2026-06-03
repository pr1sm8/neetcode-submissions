import "slices"
func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    pres := make(map[[3]int]struct{})
    res := make([][]int, 0)
    for i := range nums {
        target := -1 * nums[i]

        p0 := i+1
        p1 := len(nums) - 1
        for p0 < p1 {
            sum := nums[p0] + nums[p1]
            if sum == target {
                hashkey := [3]int{nums[i], nums[p0], nums[p1]}
                if _, ok := pres[hashkey]; !ok {
                    res = append(res, []int{ nums[i], nums[p0], nums[p1] })
                    pres[hashkey] = struct{}{}
                }
                p0 ++
            } else if sum < target {
                p0 ++
            } else {
                p1 --
            }
        }
    }
    return res
}
