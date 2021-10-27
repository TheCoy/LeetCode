package golang


var pre,head *BinaryNode

func tree2DoubleList(root *BinaryNode) *BinaryNode {
	if root == nil {
		return nil
	}
	bDfs(root)
	head.Left = pre
	pre.Right = head
	return head

}

func bDfs(root *BinaryNode) {
	if root == nil {
		return
	}
	bDfs(root.Left)
	if pre == nil {
		head = root
	}else{
		pre.Right = root
	}
	root.Left = pre
	pre = root
	bDfs(root.Right)
}