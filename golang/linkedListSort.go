package golang

func SortList(head *ListNode) *ListNode {
	return nil
}

//交换值
func quickSortV1(head, tail *ListNode) *ListNode{
	if head == tail || head.Next == tail{
		return head
	}
	cur, left := head.Next, head
	pivot := left.Val
	for cur != tail {
		if cur.Val < pivot {
			left = left.Next
			left.Val, cur.Val = cur.Val, left.Val
		}
		cur = cur.Next
	}
	head.Val, left.Val = left.Val, head.Val
	quickSortV1(head, left)
	quickSortV1(left.Next, tail)

	return head
}

//交换node
func quickSortV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	small, large := &ListNode{}, &ListNode{}
	smallHead, largeHead := small, large
	x := head.Val
	cur := head
	for cur != nil  {
		if cur.Val < x {
			small.Next = cur
			small = small.Next
		}else {
			large.Next = cur
			large = large.Next
		}
		cur = cur.Next
	}
	large.Next = nil
	small.Next = largeHead.Next

	right := quickSortV2(head.Next)
	head.Next = nil
	left := quickSortV2(smallHead.Next)

	now := left
	for now.Next != nil {
		now = now.Next
	}
	now.Next = right

	return left
}

//分隔链表
func partitionV2(head *ListNode, x int) *ListNode {
	small, large := &ListNode{}, &ListNode{}
	smallHead := small
	largeHead := large

	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		}else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}

	large.Next = nil
	small.Next = largeHead.Next

	return smallHead.Next
}
