package golang

func FindMin(nums []int) int {
	for left, right := 0, len(nums)-1; left <= right; {
		if left == right {
			return nums[left]
		}
		mid := (left + right) / 2
		if nums[left] < nums[right] {
			if nums[mid] >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] > nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return nums[0]
}
