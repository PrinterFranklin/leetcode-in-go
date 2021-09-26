#include <stdlib.h>

struct ListNode {
	int val;
	struct ListNode *next;
};

/* Solution 1: Recursion */
struct ListNode* mergeTwoLists(struct ListNode* l1, struct ListNode* l2) {
	if (l1 == NULL) {
		return l2;
	}
	if (l2 == NULL) {
		return l1;
	}
	if (l1->val < l2->val) {
		l1->next = mergeTwoLists(l1->next, l2);
		return l1;
	} else {
		l2->next = mergeTwoLists(l1, l2->next);
		return l2;
	}
}

/* Solution 2: Iteration */
struct ListNode* mergeTwoLists(struct ListNode* l1, struct ListNode* l2) {
	struct ListNode dummy;
	struct ListNode *p = &dummy;
	while (l1 != NULL && l2 != NULL) {
		if (l1->val <= l2->val) {
			p->next = l1;
			l1 = l1->next;
		} else {
			p->next = l2;
			l2 = l2->next;
		}
		p = p->next;
	}
	if (l1 == NULL) {
		p->next = l2;
	}
	if (l2 == NULL) {
		p->next = l1;
	}
	return dummy.next;
}