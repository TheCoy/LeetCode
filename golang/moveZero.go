package golang

func MoveZero(nums []int) {
    i, k := 0, 0
    for i < len(nums) {
        if nums[i] != 0 {
            nums[k], nums[i] = nums[i], nums[k]
            k++
        }
        i++
    }
}
