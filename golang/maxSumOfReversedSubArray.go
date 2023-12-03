package golang

//题意
//对于一个子数组求其最大子数组和，现在允许你将数组中的某一段翻转一次，
//那么翻转后的数组也有一个最大子数组和。
//问原数组，允许翻转一次能够得到的最大子数组和是多少？

//样例
//例如3，6，-2，-4，8，5
// 将前四个反转，得到-4，-2，6，3，8，5，最大子数组和为22（6，3，8，5），
// 也可以将后四个反转得到3，6，5，8，-4，-2，最大字数和也是22
// 其他反转方式结果都不大于22；

func maxSumOfReversedSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	dp[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		dp[i] = nums[i] + max(0, dp[i+1])
	}
	left := nums[0]
	leftMax := left
	ans := dp[0]

	for i := 1; i < n; i++ {
		ans = max(ans, leftMax+dp[i])
		left = nums[i] + max(0, left)
		leftMax = max(leftMax, left)
	}

	return ans
}