package golang

func GenNextMaxNumber(x int, nums []int) int {
	ans := 0
	digits := make([]int, 0)
	for i := x; i > 0; {
		digits = append(digits, i%10)
		i /= 10
	}

	var fn func(int, int)
	fn = func(count int, number int) {
		if count >= len(digits) {
			if number < x  {
				ans = max(ans, number)
			}else {
				ans = max(ans, number/10)
			}
			return
		}
		for i := 0; i < len(nums); i++ {
			fn(count+1, number*10+nums[i])
		}
	}
	fn(0, 0)

	return ans
}
