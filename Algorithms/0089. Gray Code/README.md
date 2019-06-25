# [89. Gray Code](https://leetcode.com/problems/gray-code/)

## 题目

The gray code is a binary numeral system where two successive values differ in only one bit.

Given a non-negative integer *n* representing the total number of bits in the code, print the sequence of gray code. A gray code sequence must begin with 0.

**Example 1:**

    Input: 2
    Output: [0,1,3,2]
    Explanation:
    00 - 0
    01 - 1
    11 - 3
    10 - 2
    
    For a given n, a gray code sequence may not be uniquely defined.
    For example, [0,2,3,1] is also a valid gray code sequence.
    
    00 - 0
    10 - 2
    11 - 3
    01 - 1

**Example 2:**

    Input: 0
    Output: [0]
    Explanation: We define the gray code sequence to begin with 0.
                 A gray code sequence of n has size = 2n, which for n = 0 the size is 20 = 1.
                 Therefore, for n = 0 the gray code sequence is [0].


## 题目大意

格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。格雷编码序列必须以 0 开头。



## 解题思路

- 输出 n 位格雷码
- 格雷码生成规则：以二进制为0值的格雷码为第零项，第一次改变最右边的位元，第二次改变右起第一个为1的位元的左边位元，第三、四次方法同第一、二次，如此反复，即可排列出 n 个位元的格雷码。
- 可以直接模拟，也可以用递归求解。
