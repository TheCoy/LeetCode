package golang

import (
    "fmt"
    "sync"
    "time"
)

func ArraySort(nums []int) {
    var wg sync.WaitGroup
    for _, n := range nums{
        wg.Add(1)
        go func(n int) {
            defer wg.Done()
            t := time.Duration(n)
            time.Sleep(t*time.Millisecond)
            fmt.Println(n)
        }(n)
    }
    wg.Wait()
}
