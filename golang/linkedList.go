package golang

type LinkedList struct {
	Next *LinkedList
	Val interface{}
}


func ReverseLinkedList(head *LinkedList) *LinkedList {
	if head == nil {
		return nil
	}
	var pre *LinkedList
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func SwapPairs(head *LinkedList) *LinkedList {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next

	head.Next = SwapPairs(newHead.Next)
	newHead.Next = head

	return newHead
}