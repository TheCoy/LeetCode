package golang

import "fmt"

type TreeNode struct {
    Left, Right *TreeNode
    Val interface{}
}

func LevelVisit(root *TreeNode) {
    stack := []*TreeNode{root}

    for len(stack) > 0 {
        node := stack[0]
        stack = stack[1:]
        fmt.Print(node.Val)
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
    }

}
