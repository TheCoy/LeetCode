/*
 * @lc app=leetcode.cn id=746 lang=golang
 * @lcpr version=30204
 *
 * [746] 使用最小花费爬楼梯
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

/*
// @lcpr case=start
// [10,15,20]\n
// @lcpr case=end

// @lcpr case=start
// [1,100,1,1,1,100,1,1,100,1]\n
// @lcpr case=end

*/

