package golang

import "container/heap"

type IHeap [][2]int

func (h IHeap) Len() int {
    return len(h)
}

func (h IHeap) Less(i,j int) bool {
    return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i,j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
    *h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
    old := *h
    res := old[len(old)-1]
    *h = old[0:len(old)-1]
    return res
}

func TopKFrequent(nums []int, k int) []int {
    numMap := make(map[int]int)
    for _, n := range nums {
        numMap[n]++
    }

    h := &IHeap{}
    heap.Init(h)
    for num, count := range numMap {
        heap.Push(h, [2]int{num, count})
        if h.Len() > k {
            h.Pop()
        }
    }

    result := make([]int, k)
    for i := 0; i < k; i++{
        result[k-i-1]  = h.Pop().([2]int)[0]
    }

    return result
}