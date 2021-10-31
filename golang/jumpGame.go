package golang


//示例 1:
//输入: nums = [2,3,1,1,4]
//输出: 2
//
//
//示例 2:
//输入: nums = [2,3,0,1,4]
//输出: 2
func JumpGame(nums []int) int {
    result := 0
    start, end := 0, 0
    for end < len(nums) - 1 {
        position := 0

        for i := start; i <= end; i++ {
            position = max(position, i+nums[i])
        }
        start = end+1
        result++
        end = position
    }

    return result
}

