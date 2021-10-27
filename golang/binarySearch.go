package golang

// 非递归版本的二分查找
func BinarySearch(target int, nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

// 递归版本的二分查找
func BinarySearchIntverse(low, high, target int, nums []int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)>>1
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return BinarySearchIntverse(mid+1, high, target, nums)
	} else {
		return BinarySearchIntverse(low, mid-1, target, nums)
	}
}

// 查找最后一个等于目标值的元素
func BinarySearchLastOne(target int, nums []int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if high > 0 && nums[high] == target {
		return high
	} else {
		return -1
	}
}
