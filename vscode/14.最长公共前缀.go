/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (40.47%)
 * Likes:    1766
 * Dislikes: 0
 * Total Accepted:    607.6K
 * Total Submissions: 1.5M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：strs = ["flower","flow","flight"]
 * 输出："fl"
 *
 *
 * 示例 2：
 *
 *
 * 输入：strs = ["dog","racecar","car"]
 * 输出：""
 * 解释：输入不存在公共前缀。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= strs.length <= 200
 * 0 <= strs[i].length <= 200
 * strs[i] 仅由小写英文字母组成
 *
 *
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	first := strs[0]
	ans := ""
	for index := 0; index < len(first); index++ {
		for i := 1; i < len(strs); i++ {
			if index >= len(strs[i]) || strs[i][index] != first[index] {
				return ans
			}
		}
		ans += string(first[index])
	}

	return ans
}

// @lc code=end

