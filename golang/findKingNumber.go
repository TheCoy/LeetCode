package golang

import (
    "math"
)

func FindKingNumber(nums []int) ([]int, error) {
    leftMax := make([]int, len(nums))
    rightMin := make([]int, len(nums))
    result := make([]int, 0)

    leftMax[0] = math.MinInt32
    for i := 1; i < len(nums); i++ {
        leftMax[i] = max(nums[i-1], leftMax[i-1])
    }
    rightMin[len(nums)-1] = math.MaxInt32
    for i := len(nums) - 2; i >= 0; i-- {
        rightMin[i] = min(nums[i+1], rightMin[i+1])
    }

    for i := 0; i < len(nums)-1; i++ {
        if nums[i] > leftMax[i] && nums[i] < rightMin[i] {
            result = append(result, nums[i])
        }
    }

    return result, nil
}

func max(x, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}

func min(x, y int) int {
    if x < y {
        return x
    } else {
        return y
    }
}
