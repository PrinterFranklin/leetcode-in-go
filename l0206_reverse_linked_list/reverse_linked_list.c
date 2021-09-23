#include <stdlib.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

/* Recursive */
struct ListNode* reverseList(struct ListNode* head) {
    if (head == NULL || head->next == NULL) {
        return head;
    }
    struct ListNode *res = reverseList(head->next);
    head->next->next = head;
    head->next = NULL;
    return res;
}

/* Iteration */
struct ListNode* reverseList2(struct ListNode* head) {
    struct ListNode *pre = NULL;
    struct ListNode *cur = head;
    while (cur) {
        struct ListNode *tmp = cur->next;
        cur->next = pre;
        pre = cur;
        cur = tmp;
    }
    return pre;
}