package golang

func ReorderOddEven(nums []int) {
	start, end := 0, len(nums)-1
	for start < end {
		if nums[start] % 2 == 1 {
			start++
		}else {
			nums[start], nums[end] = nums[end], nums[start]
			end--
		}
	}
}
