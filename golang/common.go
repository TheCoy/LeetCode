package golang


type BinaryNode struct {
	Left, Right *BinaryNode
	Val int
}


type ListNode struct {
	Val  int
	Next *ListNode
}


func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}