#include <stdlib.h>

struct ListNode {
	int val;
	struct ListNode *next;
};

struct ListNode* reverseList(struct ListNode* head, struct ListNode* tail) {
	struct ListNode* pre = tail->next;
	struct ListNode* cur = head;
	while (cur) {
		struct ListNode* tmp = cur->next;
		cur->next = pre;
		pre = cur;
		cur = tmp;
	}
	return pre;
}

struct ListNode* reverseKGroup(struct ListNode* head, int k) {
	if (!head)
		return NULL;
	struct ListNode *next_head = head;

}