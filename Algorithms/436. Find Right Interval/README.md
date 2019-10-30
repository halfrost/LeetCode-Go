# [436. Find Right Interval](https://leetcode.com/problems/find-right-interval/)


## 题目:

Given a set of intervals, for each of the interval i, check if there exists an interval j whose start point is bigger than or equal to the end point of the interval i, which can be called that j is on the "right" of i.

For any interval i, you need to store the minimum interval j's index, which means that the interval j has the minimum start point to build the "right" relationship for interval i. If the interval j doesn't exist, store -1 for the interval i. Finally, you need output the stored value of each interval as an array.

**Note:**

1. You may assume the interval's end point is always bigger than its start point.
2. You may assume none of these intervals have the same start point.

**Example 1:**

    Input: [ [1,2] ]
    
    Output: [-1]
    
    Explanation: There is only one interval in the collection, so it outputs -1.

**Example 2:**

    Input: [ [3,4], [2,3], [1,2] ]
    
    Output: [-1, 0, 1]
    
    Explanation: There is no satisfied "right" interval for [3,4].
    For [2,3], the interval [3,4] has minimum-"right" start point;
    For [1,2], the interval [2,3] has minimum-"right" start point.

**Example 3:**

    Input: [ [1,4], [2,3], [3,4] ]
    
    Output: [-1, 2, -1]
    
    Explanation: There is no satisfied "right" interval for [1,4] and [3,4].
    For [2,3], the interval [3,4] has minimum-"right" start point.

**NOTE:** input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.


## 题目大意

给定一组区间，对于每一个区间 i，检查是否存在一个区间 j，它的起始点大于或等于区间 i 的终点，这可以称为 j 在 i 的“右侧”。

对于任何区间，你需要存储的满足条件的区间 j 的最小索引，这意味着区间 j 有最小的起始点可以使其成为“右侧”区间。如果区间 j 不存在，则将区间 i 存储为 -1。最后，你需要输出一个值为存储的区间值的数组。

注意:

- 你可以假设区间的终点总是大于它的起始点。
- 你可以假定这些区间都不具有相同的起始点。


## 解题思路


- 给出一个 `interval` 的 数组，要求找到每个 `interval` 在它右边第一个 `interval` 的下标。A 区间在 B 区间的右边：A 区间的左边界的值大于等于 B 区间的右边界。
- 这一题很明显可以用二分搜索来解答。先将 `interval` 数组排序，然后针对每个 `interval`，用二分搜索搜索大于等于 `interval` 右边界值的 `interval`。如果找到就把下标存入最终数组中，如果没有找到，把 `-1` 存入最终数组中。
