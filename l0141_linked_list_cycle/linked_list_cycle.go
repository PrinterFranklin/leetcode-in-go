package leetcode0141

type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution 1: Hash Set
// Time: O(n)
// Space: O(n)
func hasCycle(head *ListNode) bool {
	set := map[*ListNode]struct{}{}
	for tmp := head; tmp != nil; tmp = tmp.Next {
		if _, ok := set[tmp]; ok {
			return true
		}
		set[tmp] = struct{}{}
	}
	return false
}

// Solution 2: Two Pointers (fast and slow)
// Time: O(n)
// Space: O(1)
func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}
