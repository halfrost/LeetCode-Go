# [923. 3Sum With Multiplicity](https://leetcode.com/problems/3sum-with-multiplicity/)

## 题目

Given an integer array A, and an integer target, return the number of tuples i, j, k  such that i < j < k and A[i] + A[j] + A[k] == target.

As the answer can be very large, return it modulo 10^9 + 7.


Example 1:

```c
Input: A = [1,1,2,2,3,3,4,4,5,5], target = 8
Output: 20
Explanation: 
Enumerating by the values (A[i], A[j], A[k]):
(1, 2, 5) occurs 8 times;
(1, 3, 4) occurs 8 times;
(2, 2, 4) occurs 2 times;
(2, 3, 3) occurs 2 times.
```


Example 2:

```c
Input: A = [1,1,2,2,2,2], target = 5
Output: 12
Explanation: 
A[i] = 1, A[j] = A[k] = 2 occurs 12 times:
We choose one 1 from [1,1] in 2 ways,
and two 2s from [2,2,2,2] in 6 ways.
```


Note:

- 3 <= A.length <= 3000
- 0 <= A[i] <= 100
- 0 <= target <= 300


## 题目大意

这道题是第 15 题的升级版。给出一个数组，要求找到 3 个数相加的和等于 target 的解组合的个数，并且要求  i < j < k。解的组合个数不需要去重，相同数值不同下标算不同解(这里也是和第 15 题的区别)

## 解题思路

这一题大体解法和第 15 题一样的，只不过算所有解组合的时候需要一点排列组合的知识，如果取 3 个一样的数，需要计算 C n 3，去 2 个相同的数字的时候，计算 C n 2，取一个数字就正常计算。最后所有解的个数都加起来就可以了。
