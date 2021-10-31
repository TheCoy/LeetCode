/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (70.14%)
 * Likes:    1027
 * Dislikes: 0
 * Total Accepted:    303.8K
 * Total Submissions: 432.4K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4]
 * 输出：[2,1,4,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = [1]
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点的数目在范围 [0, 100] 内
 * 0
 *
 *
 *
 *
 * 进阶：你能在不修改链表节点值的情况下解决这个问题吗?（也就是说，仅修改节点本身。）
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func swapPairs(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	dummy := &ListNode{
// 		Next: head,
// 	}
// 	cur := dummy.Next
// 	next := cur.Next
// 	for next != nil && cur != nil {
// 		cur.Val, next.Val = next.Val, cur.Val
// 		cur = next.Next
// 		if cur != nil {
// 			next = cur.Next
// 		} else {
// 			next = nil
// 		}
// 	}

// 	return dummy.Next
// }

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}
	prev, left, right := dummy, head, head.Next
	for left != nil && right != nil {
		tmp := right.Next
		right.Next = left
		left.Next = tmp
		prev.Next = right
		prev = left
		left = tmp
		if left == nil {
			right = nil
		} else {
			right = left.Next
		}

	}

	return dummy.Next
}

// @lc code=end

