package leetcode0160

type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution 1: Hash Set
// Time: O(m+n)
// Space: O(n)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hashTable := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		hashTable[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if hashTable[tmp] {
			return tmp
		}
	}
	return nil
}

// Solution 2: Two Pointers (Recommanded)
// Time: O(m+n)
// Space: O(1)
func getInterSectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
