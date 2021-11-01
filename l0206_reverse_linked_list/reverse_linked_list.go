package leetcode0206

type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution 1: Recursion
// Time: O(n)
// Space: O(n)
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}

// Solution 2: Iteration
// Time: O(n)
// Space: O(1)
func reverseList1(head *ListNode) *ListNode {
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
