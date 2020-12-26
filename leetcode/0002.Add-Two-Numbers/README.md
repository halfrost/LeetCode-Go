# [2. Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)

## 题目

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
```
Explanation: 342 + 465 = 807.


## 题目大意

2 个逆序的链表，要求从低位开始相加，得出结果也逆序输出，返回值是逆序结果链表的头结点。

## 解题思路

需要注意的是各种进位问题。

极端情况，例如
```
Input: (9 -> 9 -> 9 -> 9 -> 9) + (1 -> )
Output: 0 -> 0 -> 0 -> 0 -> 0 -> 1
```

为了处理方法统一，可以先建立一个虚拟头结点，这个虚拟头结点的 Next 指向真正的 head，这样 head 不需要单独处理，直接 while 循环即可。另外判断循环终止的条件不用是 p.Next ！= nil，这样最后一位还需要额外计算，循环终止条件应该是 p != nil。

c++:

```c++
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *head = nullptr, *tail = nullptr;
        int carry = 0;
        while(l1 || l2){
            int n1 = l1 ? l1->val : 0;
            int n2 = l2 ? l2->val : 0;
            int sum = n1 + n2 +carry;
            if(!head){
                head = tail = new ListNode(sum % 10);
            }else{
                tail->next = new ListNode(sum % 10);
                tail = tail->next; 
            }
            carry = sum /10;
            if(l1){
                l1 = l1->next;   
            }
            if(l2){
                l2 = l2->next;
            }
        }
        if(carry > 0){
            tail->next = new ListNode(carry);
        }
        return head;
        
    }
};
```

