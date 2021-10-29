package golang

type linkedList struct {
    Next *linkedList
    Val  int
}

func JosephRecrusion(n, m int) int {
    if n == 1 {
        return n
    }
    return (JosephRecrusion(n-1, m)+m-1)%n + 1
}

func JosephLoop(n, m int) int {
    head := createLinkedList(n)

    number := 0
    var pre *linkedList
    for head.Next != head {
        number++
        if number == m {
            pre.Next = head.Next
            head = head.Next
            number = 0
        } else {
            pre = head
            head = head.Next
        }
    }

    return head.Val
}

func createLinkedList(n int) *linkedList {
    head := &linkedList{
        Val: 1,
    }
    var cur *linkedList

    cur = head
    for i := 2; i <= n; i++ {
        new := &linkedList{
            Next: nil,
            Val:  i,
        }
        cur.Next = new
        cur = new
    }

    cur.Next = head

    return head
}
