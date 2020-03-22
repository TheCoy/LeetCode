/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	ret := make(map[int]int)
	for k, v := range nums {
		if x, ok := ret[target-v]; ok {
			return  []int{x, k}
		}
		ret[v] = k
	}
	return nil
}
// @lc code=end

