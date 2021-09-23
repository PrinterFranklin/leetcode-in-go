package leetcode0206

type ListNode struct {
	Val  int
	Next *ListNode
}

// Recursive
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}

// Iteration
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
