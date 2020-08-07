# [347. Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/)

## 题目

Given a non-empty array of integers, return the k most frequent elements.

Example 1:

```c
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
```

Example 2:

```c
Input: nums = [1], k = 1
Output: [1]
```

Note:  

- You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
- Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
 

## 题目大意

给一个非空的数组，输出前 K 个频率最高的元素。

## 解题思路

这一题是考察优先队列的题目。把数组构造成一个优先队列，输出前 K 个即可。


