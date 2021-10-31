/*
 * @lc app=leetcode.cn id=989 lang=golang
 *
 * [989] 数组形式的整数加法
 *
 * https://leetcode-cn.com/problems/add-to-array-form-of-integer/description/
 *
 * algorithms
 * Easy (47.35%)
 * Likes:    160
 * Dislikes: 0
 * Total Accepted:    50.9K
 * Total Submissions: 107.7K
 * Testcase Example:  '[1,2,0,0]\n34'
 *
 * 对于非负整数 X 而言，X 的数组形式是每位数字按从左到右的顺序形成的数组。例如，如果 X = 1231，那么其数组形式为 [1,2,3,1]。
 *
 * 给定非负整数 X 的数组形式 A，返回整数 X+K 的数组形式。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：A = [1,2,0,0], K = 34
 * 输出：[1,2,3,4]
 * 解释：1200 + 34 = 1234
 *
 *
 * 示例 2：
 *
 * 输入：A = [2,7,4], K = 181
 * 输出：[4,5,5]
 * 解释：274 + 181 = 455
 *
 *
 * 示例 3：
 *
 * 输入：A = [2,1,5], K = 806
 * 输出：[1,0,2,1]
 * 解释：215 + 806 = 1021
 *
 *
 * 示例 4：
 *
 * 输入：A = [9,9,9,9,9,9,9,9,9,9], K = 1
 * 输出：[1,0,0,0,0,0,0,0,0,0,0]
 * 解释：9999999999 + 1 = 10000000000
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= A.length <= 10000
 * 0 <= A[i] <= 9
 * 0 <= K <= 10000
 * 如果 A.length > 1，那么 A[0] != 0
 *
 *
 */

// @lc code=start
func addToArrayForm(nums []int, k int) []int {
	lenN := len(nums)
	index := 0
	result := make([]int, 0)
	for k > 0 {
		x := k % 10
		var temp int
		if index <= lenN-1 {
			temp = x + nums[lenN-1-index]

		} else {
			temp = x
		}
		result = append(result, temp)
		k /= 10
		index++
	}

	for index < len(nums) {
		result = append(result, nums[lenN-1-index])
		index++
	}

	for i := 0; i < len(result); i++ {
		if result[i] > 9 {
			if i < len(result)-1 {
				result[i+1] += result[i] / 10
			} else {
				result = append(result, result[i]/10)
			}
			result[i] = result[i] % 10
		}
	}

	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}

	return result
}

// @lc code=end

