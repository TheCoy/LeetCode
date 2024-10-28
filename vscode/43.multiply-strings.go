/*
 * @lc app=leetcode.cn id=43 lang=golang
 * @lcpr version=30204
 *
 * [43] 字符串相乘
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	result := make([]byte, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			sum := (num1[i] - '0') * (num2[j] - '0')
			result[i+j+1] += sum
			if result[i+j+1] > 9 {
				result[i+j] += result[i+j+1] / 10
				result[i+j+1] %= 10
			}
		}
	}

	if result[0] == 0 {
		result = result[1:]
	}

	for i := range result {
		result[i] += '0'
	}

	return string(result)
}

// @lc code=end

/*
// @lcpr case=start
// "2"\n"3"\n
// @lcpr case=end

// @lcpr case=start
// "123"\n"456"\n
// @lcpr case=end

*/

