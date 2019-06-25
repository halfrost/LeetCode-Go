# [845. Longest Mountain in Array](https://leetcode.com/problems/longest-mountain-in-array/)

## 题目

Let's call any (contiguous) subarray B (of A) a mountain if the following properties hold:  

- B.length >= 3
- There exists some 0 < i < B.length - 1 such that B[0] < B[1] < ... B[i-1] < B[i] > B[i+1] > ... > B[B.length - 1]
(Note that B could be any subarray of A, including the entire array A.)

Given an array A of integers, return the length of the longest mountain. 

Return 0 if there is no mountain.




Example 1:

```c
Input: [2,1,4,7,3,2,5]
Output: 5
Explanation: The largest mountain is [1,4,7,3,2] which has length 5.
```

Example 2:

```c
Input: [2,2,2]
Output: 0
Explanation: There is no mountain.
```

Note:

- 0 <= A.length <= 10000
- 0 <= A[i] <= 10000


Follow up:

- Can you solve it using only one pass?
- Can you solve it in O(1) space?

## 题目大意

这道题考察的是滑动窗口的问题。

给出一个数组，要求求出这个数组里面“山”最长的长度。“山”的意思是，从一个数开始逐渐上升，到顶以后，逐渐下降。

## 解题思路

这道题解题思路也是滑动窗口，只不过在滑动的过程中多判断一个上升和下降的状态即可。

