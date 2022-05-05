package golang

import (
    "fmt"
    "math"
)

type MyIntHeap struct {
    arrVal []int
}

func (h *MyIntHeap) insert(num int) {
    h.arrVal = append(h.arrVal, num)
    i := len(h.arrVal) - 1
    for ; h.arrVal[i/2] < num; i /= 2 {
        h.arrVal[i] = h.arrVal[i/2]
    }

    h.arrVal[i] = num
}

func (h *MyIntHeap) pop() (int, error) {
    if len(h.arrVal) <= 1 {
        return -1, fmt.Errorf("no number")
    }

    ans := h.arrVal[1]
    last := h.arrVal[len(h.arrVal)-1]
    i := 1
    for ; i < len(h.arrVal)/2; {
        cIndex := i * 2
        if cIndex < len(h.arrVal) && h.arrVal[cIndex+1] > h.arrVal[cIndex] {
            cIndex++
        }
        if last <= h.arrVal[cIndex] {
            h.arrVal[i] = h.arrVal[cIndex]
        } else {
            break
        }
        i = cIndex

    }
    h.arrVal[i] = last
    h.arrVal = h.arrVal[:len(h.arrVal)-1]

    return ans, nil
}

func TopK(input []int, k int) []int {
    result := make([]int, k)
    h := &MyIntHeap{
        arrVal: make([]int, 0),
    }
    h.arrVal = append(h.arrVal, math.MaxInt32)
    for _, n := range input {
        h.insert(n)
    }
    fmt.Println("after insert", h.arrVal)
    for i := 0; i < k; i++ {
        num, _ := h.pop()
        fmt.Printf("%dth:%+v\n", i, h.arrVal)
        result[i] = num
    }

    return result
}

