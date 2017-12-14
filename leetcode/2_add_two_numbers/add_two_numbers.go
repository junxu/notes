package add_two_numbers


type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p := l1
	q := l2
	head := new(ListNode)
	curr := head
	carry := 0

	for p != nil || q != nil {
		x, y := 0, 0
		if p != nil {
			x = p.Val
			p = p.Next
		}
		if q != nil {
			y = q.Val
			q = q.Next
		}
		curr.Next = &ListNode{(x + y + carry) % 10, nil}
		curr = curr.Next
		carry = (x + y + carry) / 10
	}

	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}

	return head.Next
}
