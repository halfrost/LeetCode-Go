# [480. Sliding Window Median](https://leetcode.com/problems/sliding-window-median/)



## 题目

Median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value. So the median is the mean of the two middle value.

Examples:  

`[2,3,4]` , the median is `3`

`[2,3]`, the median is `(2 + 3) / 2 = 2.5`

Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Your job is to output the median array for each window in the original array.

For example,  

Given nums = `[1,3,-1,-3,5,3,6,7]`, and k = 3.

    Window position                Median
    ---------------               -----
    [1  3  -1] -3  5  3  6  7       1
     1 [3  -1  -3] 5  3  6  7       -1
     1  3 [-1  -3  5] 3  6  7       -1
     1  3  -1 [-3  5  3] 6  7       3
     1  3  -1  -3 [5  3  6] 7       5
     1  3  -1  -3  5 [3  6  7]      6

Therefore, return the median sliding window as `[1,-1,-1,3,5,6]`.

**Note:**   

You may assume `k` is always valid, ie: `k` is always smaller than input array's size for non-empty array.


## 题目大意

中位数是有序序列最中间的那个数。如果序列的大小是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。

例如：

[2,3,4]，中位数是 3

[2,3]，中位数是 (2 + 3) / 2 = 2.5

给出一个数组 nums，有一个大小为 k 的窗口从最左端滑动到最右端。窗口中有 k 个数，每次窗口移动 1 位。你的任务是找出每次窗口移动后得到的新窗口中元素的中位数，并输出由它们组成的数组。



## 解题思路


- 给定一个数组和一个窗口为 K 的窗口，当窗口从数组的左边滑动到数组右边的时候，输出每次移动窗口以后，在窗口内的排序之后中间大小的值。
- 这一题是第 239 题的升级版。
- 这道题最暴力的方法就是将窗口内的元素都排序，时间复杂度 O(n * K)。
- 另一种思路是用两个优先队列，大顶堆里面的元素都比小顶堆里面的元素小。小顶堆里面存储排序以后中间靠后的值大的元素，大顶堆里面存储排序以后中间靠前的值小的元素。如果 k 是偶数，那么两个堆都有 k/2 个元素，中间值就是两个堆顶的元素；如果 k 是奇数，那么小顶堆比大顶堆多一个元素，中间值就是小顶堆的堆顶元素。删除一个元素，元素都标记到删除的堆中，取 top 的时候注意需要取出没有删除的元素。时间复杂度 O(n * log k) 空间复杂度 O(k)
