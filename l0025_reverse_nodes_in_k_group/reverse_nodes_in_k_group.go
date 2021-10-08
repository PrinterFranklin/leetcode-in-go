package leetcode0025

// ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next
	cur := head
	for pre != tail {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return tail, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		tmp := tail.Next
		head, tail = reverseList(head, tail)
		pre.Next = head
		tail.Next = tmp
		pre = tail
		head = tail.Next
	}
	return dummy.Next
}
