/*
 * @lc app=leetcode.cn id=2639 lang=golang
 * @lcpr version=30122
 *
 * [2639] 查询网格图中每一列的宽度
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findColumnWidth(grid [][]int) []int {
	ans := make([]int, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			ans[i] = max(ans[i], lenOfInt(grid[j][i]))
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lenOfInt(a int) int {

	return len(strconv.Itoa(a))
}

// @lc code=end

/*
// @lcpr case=start
// [[1],[22],[333]]\n
// @lcpr case=end

// @lcpr case=start
// [[-15,1,3],[15,7,12],[5,6,-2]]\n
// @lcpr case=end

*/

