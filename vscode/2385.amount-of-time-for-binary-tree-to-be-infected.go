/*
 * @lc app=leetcode.cn id=2385 lang=golang
 * @lcpr version=30122
 *
 * [2385] 感染二叉树需要的总时间
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func amountOfTime(root *TreeNode, start int) int {
	graph := make(map[int][]int, 0)

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		for _, child := range []*TreeNode{node.Left, node.Right} {
			if child != nil {
				graph[node.Val] = append(graph[node.Val], child.Val)
				graph[child.Val] = append(graph[child.Val], node.Val)
				dfs(child)
			}
		}
	}
	dfs(root)

	visited := make(map[int]bool, 0)
	queue := [][]int{{start, 0}}
	time := 0
	visited[start] = true

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		val := head[0]
		time = head[1]
		for _, child := range graph[val] {
			if !visited[child] {
				visited[child] = true
				queue = append(queue, []int{child, time + 1})
			}
		}
	}

	return time
}

// @lc code=end

/*
// @lcpr case=start
// [1,5,3,null,4,10,6,9,2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

*/

