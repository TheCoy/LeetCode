/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (33.01%)
 * Likes:    3720
 * Dislikes: 0
 * Total Accepted:    633.4K
 * Total Submissions: 1.9M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0
 * 且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-1,0,1,2,-1,-4]
 * 输出：[[-1,-1,2],[-1,0,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [0]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * -10^5
 *
 *
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && (nums[i] == nums[i-1]) {
			continue
		}
		right := len(nums) - 1
		for left := i + 1; left < len(nums); left++ {
			if left > i+1 && (nums[left] == nums[left-1]) {
				continue
			}
			for ; left < right && (nums[left]+nums[right]+nums[i] > 0); right-- {

			}
			if left == right {
				break
			}
			if nums[left]+nums[right]+nums[i] == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
			}
		}
	}

	return ans
}

// @lc code=end

