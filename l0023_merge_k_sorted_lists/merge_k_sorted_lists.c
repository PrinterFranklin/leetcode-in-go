#include <stdlib.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

struct ListNode* mergeTwoLists(struct ListNode *l1, struct ListNode *l2) {
    if (l1 == NULL)
        return l2;
    if (l2 == NULL)
        return l1;
    if (l1->val < l2->val) {
        l1->next = mergeTwoLists(l1->next, l2);
        return l1;
    } else {
        l2->next = mergeTwoLists(l1, l2->next);
        return l2;
    }
}

/* a hepler function */
struct ListNode* mergeLists(struct ListNode** lists, int left, int right) {
    if (left == right)
        return lists[left];
    if (left > right)
        return NULL;
    int mid = left + (right - left) / 2;
    struct ListNode *l1 = mergeLists(lists, left, mid);
    struct ListNode *l2 = mergeLists(lists, mid+1, right);
    return mergeTwoLists(l1, l2);
}

struct ListNode* mergeKLists(struct ListNode** lists, int listsSize) {
    if (!listsSize)
        return NULL;
    return mergeLists(lists, 0, listsSize-1);
}