package golang

func createTargetArray(nums []int, index []int) []int {
	result := make([]int, len(nums))
	for i, idx := range index {
		copy(result[idx+1:], result[idx:])
		result[i] = nums[i]
	}

	return result
}
