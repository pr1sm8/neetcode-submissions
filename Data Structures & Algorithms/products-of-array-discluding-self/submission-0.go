func productExceptSelf(nums []int) []int {
    prefixProduct := make([]int, len(nums))
    suffixProduct := make([]int, len(nums))
    
    cur := 1
    for i, val := range nums {
        cur *= val
        prefixProduct[i] = cur
    }

    // fmt.Println(prefixProduct)
    cur = 1
    for i:=len(nums)-1; i > -1 ; i -- {
        cur *= nums[i]
        suffixProduct[i] = cur
    }
    // fmt.Println(suffixProduct)
    res := make([]int, len(nums))
    for i := range res {
        left := i - 1
        right := i + 1
        prefix := 1
        if left >= 0 {
            prefix = prefixProduct[left]
        }
        suffix := 1
        if right < len(nums) {
            suffix = suffixProduct[right]
        }
        res[i] = prefix * suffix
    }
    return res
}
