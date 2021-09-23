#include <stdlib.h>

struct ListNode {
	int var;
	struct ListNode *next;
};

struct ListNode* removeNthFromEnd(struct ListNode* head, int n) {
	struct ListNode *dummy = malloc(sizeof(struct ListNode));
	dummy->next = head;
	struct ListNode *fast = dummy;
	struct ListNode *slow = dummy;
	for (int i = 0; i < n; i++) {
		fast = fast->next;
	}
	while (fast->next) {
		fast = fast->next;
		slow = slow->next;
	}
	slow->next = slow->next->next;
	return dummy->next;
}