# [307. Range Sum Query - Mutable](https://leetcode.com/problems/range-sum-query-mutable/)


## 题目:

Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

The update(i, val) function modifies nums by updating the element at index i to val.

**Example:**

    Given nums = [1, 3, 5]
    
    sumRange(0, 2) -> 9
    update(1, 2)
    sumRange(0, 2) -> 8

**Note:**

1. The array is only modifiable by the update function.
2. You may assume the number of calls to update and sumRange function is distributed evenly.


## 题目大意

给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。

update(i, val) 函数可以通过将下标为 i 的数值更新为 val，从而对数列进行修改。

示例:

```
Given nums = [1, 3, 5]

sumRange(0, 2) -> 9  
update(1, 2)  
sumRange(0, 2) -> 8  
```

说明:

- 数组仅可以在 update 函数下进行修改。
- 你可以假设 update 函数与 sumRange 函数的调用次数是均匀分布的。

## 解题思路


- 给出一个数组，数组里面的数都是`**可变**`的，设计一个数据结构能够满足查询数组任意区间内元素的和。
- 对比第 303 题，这一题由于数组里面的元素都是**`可变`**的，所以第一个想到的解法就是线段树，构建一颗线段树，父结点内存的是两个子结点的和，初始化建树的时间复杂度是 O(log n)，查询区间元素和的时间复杂度是 O(log n)，更新元素值的时间复杂度是 O(log n)。
- 如果此题还用 prefixSum 的思路解答呢？那每次 update 操作的时间复杂度都是 O(n)，因为每次更改一个值，最坏情况就是所有的 prefixSum 都要更新一次。prefixSum 的方法在这道题上面也可以 AC，只不过时间排名在 5%，非常差。
