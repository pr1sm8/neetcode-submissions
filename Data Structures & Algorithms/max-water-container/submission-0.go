func maxArea(heights []int) int {
    p0 := 0
    p1 := len(heights) - 1

    maxCapacity := 0
    for p0 < p1 {
        minHeight := min(heights[p0], heights[p1])
        capacity := minHeight * (p1 - p0)
        maxCapacity = max(capacity, maxCapacity)
        if heights[p0] <= heights[p1] {
            p0 ++
        } else {
            p1 --
        }
    }
    return maxCapacity
}
