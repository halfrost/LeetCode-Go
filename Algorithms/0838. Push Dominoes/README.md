# [838. Push Dominoes](https://leetcode.com/problems/push-dominoes/)

## 题目

There are N dominoes in a line, and we place each domino vertically upright.

In the beginning, we simultaneously push some of the dominoes either to the left or to the right.

![](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/05/18/domino.png)


After each second, each domino that is falling to the left pushes the adjacent domino on the left.

Similarly, the dominoes falling to the right push their adjacent dominoes standing on the right.

When a vertical domino has dominoes falling on it from both sides, it stays still due to the balance of the forces.

For the purposes of this question, we will consider that a falling domino expends no additional force to a falling or already fallen domino.

Given a string "S" representing the initial state. S[i] = 'L', if the i-th domino has been pushed to the left; S[i] = 'R', if the i-th domino has been pushed to the right; S[i] = '.', if the i-th domino has not been pushed.

Return a string representing the final state. 


Example 1:

```c
Input: ".L.R...LR..L.."
Output: "LL.RR.LLRRLL.."
```

Example 2:

```c
Input: "RR.L"
Output: "RR.L"
Explanation: The first domino expends no additional force on the second domino.
```


Note:

- 0 <= N <= 10^5
- String dominoes contains only 'L', 'R' and '.'


## 题目大意

这道题是一个道模拟题，考察的也是滑动窗口的问题。

给出一个字符串，L 代表这个多米诺骨牌会往左边倒，R 代表这个多米诺骨牌会往右边倒，问最终这些牌倒下去以后，情况是如何的，输出最终情况的字符串。

## 解题思路

这道题可以预先在初始字符串头和尾都添加一个字符串，左边添加 L，右边添加 R，辅助判断。

