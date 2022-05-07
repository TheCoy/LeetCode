package golang

func FindDuplicate(nums []int) int {
	bitMax := 1
	ans := 0

	n := len(nums)
	for ((n-1) >> uint(bitMax)) != 0 {
		bitMax++
	}
	for bit := 0; bit <= bitMax; bit++ {
		x, y := 0, 0
		for i := 0; i < len(nums); i++ {
			if ((nums[i] >> uint(bit)) & 1) > 0 {
				x++
			}
			if i >= 1 && (i >> uint(bit) & 1) > 0 {
				y++
			}
		}
		if x > y {
			ans |= (1  << uint(bit))
		}
	}
	return ans
}

