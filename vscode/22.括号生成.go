/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (77.25%)
 * Likes:    1948
 * Dislikes: 0
 * Total Accepted:    319.2K
 * Total Submissions: 413.2K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：["((()))","(()())","(())()","()(())","()()()"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：["()"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 */

// @lc code=start
func generateParenthesis(n int) []string {
	result := []string{}

	var loop func(left, right int, path string)
	loop = func(left, right int, path string) {
		if len(path) == 2*n {
			result = append(result, path)
		}

		if left > 0 {
			loop(left-1, right, path+"(")
		}

		if left < right {
			loop(left, right-1, path+")")
		}
	}

	loop(n, n, "")
	return result
}

// @lc code=end

