package golang

import (
    "container/heap"
    "fmt"
)

type IntHeap []int


var _ heap.Interface = &IntHeap{}


func (h IntHeap) Len() int {
    return len(h)
}

func (h IntHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Less(i, j int) bool {
    return h[i] > h[j]
}

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    x := old[len(old)-1]
    *h = old[0:len(old)-1]

    return x
}

func FindTopK(nums []int, k int) int {
    s := &IntHeap{}
    heap.Init(s)
    for _, n := range nums {
        heap.Push(s, n)
    }


    var result int
    for i:=0;i<k;i++{
        result = heap.Pop(s).(int)
        fmt.Print(i, ":", result, "\n")
    }

    return result
}
