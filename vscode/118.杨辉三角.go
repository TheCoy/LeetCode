/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 *
 * https://leetcode-cn.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (72.23%)
 * Likes:    563
 * Dislikes: 0
 * Total Accepted:    203.7K
 * Total Submissions: 280.7K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
 *
 * 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
 *
 *
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: numRows = 5
 * 输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
 *
 *
 * 示例 2:
 *
 *
 * 输入: numRows = 1
 * 输出: [[1]]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1
 *
 *
 */

// @lc code=start
func generate(numRows int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{1})
	for i := 1; i < numRows; i++ {
		newRow := make([]int, i+1)
		newRow[0] = 1
		newRow[i] = 1
		for j := 1; j < i; j++ {
			newRow[j] = result[i-1][j] + result[i-1][j-1]
		}
		result = append(result, newRow)
	}
	return result
}

// @lc code=end

