/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (48.30%)
 * Likes:    571
 * Dislikes: 0
 * Total Accepted:    258.9K
 * Total Submissions: 534.5K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 *
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 *
 * 说明：叶子节点是指没有子节点的节点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [2,null,3,null,4,null,5,null,6]
 * 输出：5
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数的范围在 [0, 10^5] 内
 * -1000
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var loop func(*TreeNode, int)
	result := math.MaxInt32
	loop = func(node *TreeNode, depth int) {
		if node.Left == nil && node.Right == nil {
			if depth < result {
				result = depth
			}
			return
		}
		if node.Left != nil {
			loop(node.Left, depth+1)
		}
		if node.Right != nil {
			loop(node.Right, depth+1)
		}
	}
	loop(root, 1)

	return result
}

// @lc code=end

