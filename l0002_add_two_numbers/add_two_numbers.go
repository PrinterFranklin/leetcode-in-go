package leetcode0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := ListNode{0, nil}
	tail := &dummy
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		tail.Next = &ListNode{Val: sum}
		tail = tail.Next
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
