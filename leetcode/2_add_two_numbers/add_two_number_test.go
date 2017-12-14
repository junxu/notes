package add_two_numbers

import (
	"testing"
	"fmt"
)

func TestAddTwoNumbers(t *testing.T) {
	list1 := [][]int{
		[]int{3, 2, 4},
		[]int{8, 7, 3 , 1},
		[]int{8, 1, 9, 9, 9},
	}
	list2 := [][]int{
		[]int{5, 8, 5},
		[]int{2, 1},
		[]int{2, 9},
	}
	result := [][]int{
		[]int{8, 0, 0, 1},
		[]int{0, 9, 3, 1},
		[]int{0, 1, 0, 0, 0, 1},
	}

	for i := 0; i < 3; i++ {
		l1, l2 := generateList(list1[i]), generateList(list2[i])
		//l1 := generateList(list1[i])
		//l2 := generateList(list2[i])
		ret := AddTwoNumbers(l1, l2)
		printList(l1)
		printList(l2)
		printList(ret)
		printList(generateList(result[i]))
		if !equal(ret, generateList(result[i])) {
			t.Fatalf("case %d failes: %v\n", i, ret.Val)
		}
	}
}

func generateList(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}

	head := new(ListNode)
	curr := head
	for i, val := range list {
		curr.Val = val
		if i != len(list) -1 {
			curr.Next = new(ListNode)
			curr = curr.Next
		}
	}

	return head
}

func equal(l1, l2 *ListNode) bool {
	p, q := l1, l2
	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		p, q = p.Next, q.Next
	}

	if p != nil || q != nil {
		return false
	}
	return true
}

func printList(l *ListNode) {
	cur := l
	if cur != nil {
		fmt.Printf("%d", cur.Val)
	}
	cur = l.Next
	for cur != nil {
		fmt.Printf(" -> %d", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("\n")
}
