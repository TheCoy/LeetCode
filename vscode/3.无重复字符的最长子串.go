/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (34.68%)
 * Likes:    3758
 * Dislikes: 0
 * Total Accepted:    513.8K
 * Total Submissions: 1.5M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 * 
 * 示例 1:
 * 
 * 输入: "abcabcbb"
 * 输出: 3 
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 * 
 * 
 * 示例 2:
 * 
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 * 
 * 
 * 示例 3:
 * 
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 * 
 * 
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	hashMap := make(map[byte]int)
	right, max, length := -1, 0, len(s)

	for i := 0; i < length; i++ {
		if i != 0 {
			delete(hashMap, s[i - 1])
		}

		for right + 1 < length && hashMap[s[right + 1]] == 0 {
			right++
			hashMap[s[right]]++
		}

		if right - i + 1 > max {
			max = right - i + 1
		}
	}

	return max
}
// @lc code=end

