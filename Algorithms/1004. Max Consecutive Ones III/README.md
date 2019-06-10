# [1004. Max Consecutive Ones III](https://leetcode.com/problems/max-consecutive-ones-iii/)

## 题目

Given an array A of 0s and 1s, we may change up to K values from 0 to 1.

Return the length of the longest (contiguous) subarray that contains only 1s. 


Example 1:

```c
Input: A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
Output: 6
Explanation: 
[1,1,1,0,0,1,1,1,1,1,1]
Bolded numbers were flipped from 0 to 1.  The longest subarray is underlined.
```

Example 2:

```c
Input: A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
Output: 10
Explanation: 
[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1.  The longest subarray is underlined.
```


Note:

- 1 <= A.length <= 20000
- 0 <= K <= A.length
- A[i] is 0 or 1 


## 题目大意

这道题考察的是滑动窗口的问题。

给出一个数组，数组中元素只包含 0 和 1 。再给一个 K，代表能将 0 变成 1 的次数。要求出经过变换以后，1 连续的最长长度。

## 解题思路

按照滑动窗口的思路处理即可，不断的更新和维护最大长度。
