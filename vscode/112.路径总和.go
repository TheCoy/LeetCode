/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
 *
 * https://leetcode-cn.com/problems/path-sum/description/
 *
 * algorithms
 * Easy (52.54%)
 * Likes:    653
 * Dislikes: 0
 * Total Accepted:    254.4K
 * Total Submissions: 483.9K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,null,1]\n22'
 *
 * 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点
 * 的路径，这条路径上所有节点值相加等于目标和 targetSum 。
 *
 * 叶子节点 是指没有子节点的节点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,3], targetSum = 5
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1,2], targetSum = 0
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目在范围 [0, 5000] 内
 * -1000
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	result := false
	var foo func(*TreeNode, int)
	foo = func(node *TreeNode, total int) {
		if node.Left == nil && node.Right == nil {
			if total+node.Val == targetSum {
				result = true
			}
			return
		}
		if node.Left != nil {
			foo(node.Left, total+node.Val)
		}
		if node.Right != nil {
			foo(node.Right, total+node.Val)
		}
	}

	foo(root, 0)

	return result
}

// @lc code=end

