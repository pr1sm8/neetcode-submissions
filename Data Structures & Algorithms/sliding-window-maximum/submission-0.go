import (
	// "container/heap"
	// "slices"
	// "fmt"
)

type ValIndex struct {
	val int
	index int
}

type IntHeap []ValIndex

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(ValIndex))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h IntHeap) Peak() ValIndex {
	return h[0]
}

func maxSlidingWindow(nums []int, k int) []int {
    hb := &IntHeap{}
	heap.Init(hb)

	output := []int{}
	
	for i := range k {
		heap.Push(hb, ValIndex{val: nums[i], index: i})
	}

	output = append(output, hb.Peak().val)

	for i := k; i<len(nums); i++ {
		heap.Push(hb, ValIndex{val: nums[i], index: i})

		var can ValIndex
		for {
			can = hb.Peak();
			if can.index > i-k && can.index <= i {
				break
			}
			heap.Pop(hb)
		}

		output = append(output, can.val)
	}
	return output
}
