# [56. Merge Intervals](https://leetcode.com/problems/merge-intervals/)

## 题目

Given a collection of intervals, merge all overlapping intervals.

Example 1:  

```
Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
```

Example 2:  

```
Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
```

## 题目大意

合并给的多个区间，区间有重叠的要进行区间合并。


## 解题思路

先按照区间起点进行排序。然后从区间起点小的开始扫描，依次合并每个有重叠的区间。