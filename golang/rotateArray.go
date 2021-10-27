package golang

func Rotate(nums []int, k int)  {
	pivot := k % len(nums)
	Reverse(nums, 0, len(nums)-1)
	Reverse(nums, 0, pivot-1)
	Reverse(nums, pivot, len(nums)-1)
}

func Reverse(nums []int, start, end int) {
	for i:= 0;i<(end-start)/2;i++{
		nums[start+i], nums[end-i] = nums[end-i], nums[start+i]
	}
}
