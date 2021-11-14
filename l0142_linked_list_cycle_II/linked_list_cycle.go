package leetcode0142

type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution 1: hash map
// Time: O(n)
// Space: O(n)
func detectCycle(head *ListNode) *ListNode {
	track := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := track[head]; ok {
			return head
		}
		track[head] = struct{}{}
		head = head.Next
	}
	return nil
}

// Solution 2: fast and slow pointer (recommanded)
// Time: O(n)
// Space: O(1)
func detectCycle1(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
