/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (43.52%)
 * Likes:    2550
 * Dislikes: 0
 * Total Accepted:    721K
 * Total Submissions: 1.6M
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "()"
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "()[]{}"
 * 输出：true
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "(]"
 * 输出：false
 *
 *
 * 示例 4：
 *
 *
 * 输入：s = "([)]"
 * 输出：false
 *
 *
 * 示例 5：
 *
 *
 * 输入：s = "{[]}"
 * 输出：true
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 仅由括号 '()[]{}' 组成
 *
 *
 */

// @lc code=start
func isValid(s string) bool {
	stack := make([]byte, 0)
	rule := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, b := range []byte(s) {
		if _, ok := rule[b]; ok {
			stack = append(stack, b)
			continue
		}

		if len(stack) <= 0 {
			return false
		}
		d := stack[len(stack)-1]
		if rule[d] == b {
			stack = stack[0 : len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

// @lc code=end

