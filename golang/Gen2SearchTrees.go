package golang

func Generate2SearchTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateTree(1, n)
}

func generateTree(start, end int) []*TreeNode {
	if start >= end {
		return []*TreeNode{}
	}

	ans := make([]*TreeNode, 0)

	for i := start; i < end; i++ {
		leftNodes := generateTree(start, i-1)
		rightNodes := generateTree(i+1, end)

		for j := 0; j < len(leftNodes); j++ {
			for k := 0; k < len(rightNodes); k++ {
				root := &TreeNode{Val: i}
				root.Left = leftNodes[j]
				root.Right = rightNodes[k]
				ans = append(ans, root)
			}
		}
	}

	return ans
}


func InsertIntoTree(root *TreeNode, n int) *TreeNode {
	if root == nil {
		return &TreeNode{Val:n}
	}

	if n > root.Val {
		root.Right = InsertIntoTree(root.Right, n)
	}else {
		root.Left = InsertIntoTree(root.Left, n)
	}

	return root
}