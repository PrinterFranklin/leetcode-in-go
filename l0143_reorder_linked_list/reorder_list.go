package leetcode0143

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	// find the middle node
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// reverse the latter part
	l1, l2 := head, slow.Next
	slow.Next = nil
	l2 = reverse(l2)
	// merge two linked list (the latter part is always shorter)
	for l2 != nil {
		tmp := l2.Next
		l2.Next = l1.Next
		l1.Next = l2
		l1 = l2.Next
		l2 = tmp
	}
}

func reverse(head *ListNode) *ListNode {
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
