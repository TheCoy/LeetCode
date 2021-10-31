package golang

func QuickSort(arr []int, start, end int) {
    if start < end {
        pivot := partition(arr, start, end)
        QuickSort(arr, start, pivot-1)
        QuickSort(arr, pivot+1, end)
    }
}

func partition(arr []int, start, end int) int {
    m := arr[end]
    k := start
    for i := start; i < end; i++ {
        if arr[i] > m {
            arr[i], arr[k] = arr[k], arr[i]
            k++
        }
    }

    arr[k], arr[end] = arr[end], arr[k]

    return k
}

func QuickSortLoop(nums []int, start, end int) {
    if start >= end {
        return
    }
    m := nums[end]
    k := start
    for i := start; i < end; i++ {
        if nums[i] < m {
            nums[i], nums[k] = nums[k], nums[i]
            k++
        }
    }
    nums[k], nums[end] = nums[end], nums[k]
    QuickSortLoop(nums, start, k-1)
    QuickSortLoop(nums, k+1, end)
}
