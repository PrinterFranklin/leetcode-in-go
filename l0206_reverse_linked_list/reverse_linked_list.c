#include <stdlib.h>

struct ListNode {
    int var;
    struct ListNode *next;
};

struct ListNode* reverseList(struct ListNode* head) {
    if (head == NULL || head->next == NULL) {
        return head;
    }
    struct ListNode *res = reverseList(head->next);
    head->next->next = head;
    head->next = NULL;
    return res;
}