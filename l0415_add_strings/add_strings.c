#include <string.h>

char* addStrings(char* num1, char* num2) {
    int i = strlen(num1) - 1;
    int j = strlen(num2) - 1;
    int carry = 0;
    // 1 for array size, 1 for carry, 1 for '\0'
    char *res = (char*)malloc(sizeof(char) * (max(i, j) + 3));
    int len = 0;

    while (i >= 0 || j >= 0 || carry != 0) {
        int x = i >= 0 ? num1[i] - '0' : 0;
        int y = j >= 0 ? num2[j] - '0' : 0;
        int sum = x + y + carry;
        res[len++] = '0' + sum % 10;
        carry = sum / 10;
        i--;
        j--;
    }
    for (int i = 0; 2 * i < len; i++) {
        int t = res[i];
        res[i] = res[len - i - 1], res[len - i - 1] = t;
    }
    res[len++] = 0;
    return res;
}

int max(int i, int j) {
    if (i > j) {
        return i;
    }
    return j;
}