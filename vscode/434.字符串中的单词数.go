/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 *
 * https://leetcode-cn.com/problems/number-of-segments-in-a-string/description/
 *
 * algorithms
 * Easy (37.38%)
 * Likes:    95
 * Dislikes: 0
 * Total Accepted:    37.4K
 * Total Submissions: 100.1K
 * Testcase Example:  '"Hello, my name is John"'
 *
 * 统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
 *
 * 请注意，你可以假定字符串里不包括任何不可打印的字符。
 *
 * 示例:
 *
 * 输入: "Hello, my name is John"
 * 输出: 5
 * 解释: 这里的单词是指连续的不是空格的字符，所以 "Hello," 算作 1 个单词。
 *
 *
 */

// @lc code=start
func countSegments(s string) int {
	cnt := 0
	flag := true
	for _, c := range s {
		if c != ' ' && flag {
			cnt++
			flag = false
		}
		if c == ' ' {
			flag = true
		}

	}

	return cnt
}

// @lc code=end

