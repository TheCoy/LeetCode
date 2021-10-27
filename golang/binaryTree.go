package golang

func inorderTraverse(root *TreeNode) []interface{} {
	result := make([]interface{}, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil{

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, node.Val)

		root = node.Right
	}

	return result
}

func preOrderTree(root *TreeNode) []interface{} {
	result := make([]interface{}, 0)
	stack := make([]*TreeNode, 0)

	if root == nil {
		return result
	}

	for len(stack) > 0 || root != nil {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root.Right)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	return result
}