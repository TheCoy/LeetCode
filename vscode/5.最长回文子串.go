/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (34.96%)
 * Likes:    4056
 * Dislikes: 0
 * Total Accepted:    713K
 * Total Submissions: 2M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出："bb"
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "a"
 * 输出："a"
 *
 *
 * 示例 4：
 *
 *
 * 输入：s = "ac"
 * 输出："a"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 仅由数字和英文字母（大写和/或小写）组成
 *
 *
 */

// @lc code=start
func longestPalindrome(s string) string {
	var left, right int
	ans := 0
	for i := 0; i < len(s); i++ {
		l1, r1 := expand(s, i, i)
		l2, r2 := expand(s, i, i+1)
		if r1-l1+1 > ans {
			left, right = l1, r1
			ans = r1 - l1 + 1
		}
		if r2-l2+1 > ans {
			left, right = l2, r2
			ans = r2 - l2 + 1
		}
	}
	return s[left : right+1]
}

func expand(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right] && left <= right; left, right = left-1, right+1 {

	}
	return left + 1, right - 1
}

// @lc code=end

