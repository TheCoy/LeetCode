/*
 * @lc app=leetcode.cn id=2739 lang=golang
 * @lcpr version=30122
 *
 * [2739] 总行驶距离
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func distanceTraveled(mainTank int, additionalTank int) int {
	total := 0
	left := mainTank
	for ; left >= 5; left -= 5 {
		total += 5
		if additionalTank > 0 {
			left++
			additionalTank--
		}
	}
	total += left
	return total * 10
}

// @lc code=end

/*
// @lcpr case=start
// 5\n10\n
// @lcpr case=end

// @lcpr case=start
// 1\n2\n
// @lcpr case=end

*/

