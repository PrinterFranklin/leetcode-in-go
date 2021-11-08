package leetcode0025

// ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution 1: Iteration
// Time: O(n)
// Space: O(1)
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, end := dummy, head
	for end != nil {
		for i := 0; i < k-1; i++ {
			end = end.Next
			if end == nil {
				return dummy.Next
			}
		}
		start := pre.Next
		next := end.Next
		end.Next = nil
		pre.Next = reverseList(start)
		start.Next = next
		pre = start
		end = pre.Next
	}
	return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
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

// Solution 2: Recursion
// Time: O(n)
// Space: O(n)
func reverseKGroup1(head *ListNode, k int) *ListNode {
	next := head
	for i := 0; i < k; i++ {
		if next == nil {
			return head
		}
		next = next.Next
	}
	pre := reverseKGroup1(next, k)
	cur := head
	for i := 0; i < k; i++ {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
