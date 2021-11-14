package leetcode0092

type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution 1: Reuse the reverse function (traverse twice)
// Time: O(n)
// Space: O(1)
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	start, end := pre.Next, pre.Next
	for i := 0; i < right-left; i++ {
		end = end.Next
	}
	next := end.Next
	end.Next = nil
	pre.Next = reverse(start)
	start.Next = next

	return dummy.Next
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

// Solution 2: insert to head (recommended)
// Time: O(n)
// Space: O(1)
func reverseBetween1(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}
