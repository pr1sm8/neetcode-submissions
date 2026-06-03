func longestConsecutive(nums []int) int {
    starts := make(map[int]struct{})
    present := make(map[int]struct{})
    for _, num := range nums {
        present[num] = struct{}{}
    }
    for _, num := range nums {
        if _, ok := present[num-1]; !ok {
            starts[num] = struct{}{}
        }
    }
    maxSeq := 0
    for start := range starts {
        i := start+1
        seq := 1
        for {
            _, ok := present[i]
            if !ok {
                break
            }
            i++
            seq++
        }
        if seq > maxSeq {
            maxSeq = seq
        }
    }
    return maxSeq
}
