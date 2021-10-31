/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 *
 * https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (54.59%)
 * Likes:    1016
 * Dislikes: 0
 * Total Accepted:    204K
 * Total Submissions: 372.4K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left  。请你反转从位置 left 到位置 right 的链表节点，返回
 * 反转后的链表 。
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], left = 2, right = 4
 * 输出：[1,4,3,2,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [5], left = 1, right = 1
 * 输出：[5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目为 n
 * 1
 * -500
 * 1
 *
 *
 *
 *
 * 进阶： 你可以使用一趟扫描完成反转吗？
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

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	pre := dummy
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dummy.Next
}

// func reverseBetween(head *ListNode, left int, right int) *ListNode {
// 	// if head.Next == nil || head == nil {
// 	// 	return head
// 	// }
// 	// pre, next, cur := head, head, head
// 	// for i := 1; i <= right; i++ {
// 	// 	if i == left-1 {
// 	// 		pre = cur
// 	// 	}
// 	// 	if i == right {
// 	// 		next = cur
// 	// 	}
// 	// 	cur = cur.Next
// 	// }
// 	// tail := next.Next
// 	// next.Next = nil
// 	// newH := reverse(pre.Next)
// 	// pre.Next.Next = tail
// 	// pre.Next = newH

// 	return head
// }

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

// @lc code=end

