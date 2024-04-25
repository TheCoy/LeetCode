/*
 * @lc app=leetcode.cn id=39 lang=golang
 * @lcpr version=30122
 *
 * [39] 组合总和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	// sort.Ints(candidates)
	var dfs func(int, int, []int)
	dfs = func(target int, index int, path []int) {
		if target < 0 {
			return
		}
		if target == 0 {
			ans = append(ans, path)
			return
		}
		for key, val := range candidates {
			if key >= index && target >= val {
				newPath := make([]int, 0, len(path))
				newPath = append(newPath, path...)
				newPath = append(newPath, val)
				dfs(target-val, key, newPath)
			}
		}
	}
	dfs(target, 0, []int{})
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,6,7]\n7\n
// @lcpr case=end

// @lcpr case=start
// [2,3,5]\n8\n
// @lcpr case=end

// @lcpr case=start
// [2]\n1\n
// @lcpr case=end

*/

