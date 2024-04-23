package golang

func NextGreaterElements(nums []int) []int {
	ans := make([]int, len(nums))
	for k, _ := range ans {
		res := -1
		for i := k + 1; i != k; i = (i + 1) % len(nums) {
			i = i % len(nums)
			if nums[i] > nums[k] {
				res = nums[i]
				break
			}
		}
		ans[k] = res
	}

	return ans
}
