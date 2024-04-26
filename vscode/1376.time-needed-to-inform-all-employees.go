/*
 * @lc app=leetcode.cn id=1376 lang=golang
 * @lcpr version=30122
 *
 * [1376] 通知所有员工所需的时间
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	roadMap := make(map[int][]int, 0)

	for i, m := range manager {
		if _, ok := roadMap[m]; !ok {
			roadMap[m] = make([]int, 0)
		}
		roadMap[m] = append(roadMap[m], i)
	}

	var dfs func(int) int
	dfs = func(cur int) int {
		ans := 0
		for _, next := range roadMap[cur] {
			ans = max(ans, dfs(next))
		}
		return informTime[cur] + ans
	}
	return dfs(headID)

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

/*
// @lcpr case=start
// 1\n0\n[-1]\n[0]\n
// @lcpr case=end

// @lcpr case=start
// 6\n2\n[2,2,-1,2,2,2]\n[0,0,1,0,0,0]\n
// @lcpr case=end

*/

