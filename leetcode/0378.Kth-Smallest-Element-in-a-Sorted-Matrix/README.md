# [378. Kth Smallest Element in a Sorted Matrix](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/)


## 题目:

Given a n x n matrix where each of the rows and columns are sorted in ascending order, find the kth smallest element in the matrix.

Note that it is the kth smallest element in the sorted order, not the kth distinct element.

**Example:**

    matrix = [
       [ 1,  5,  9],
       [10, 11, 13],
       [12, 13, 15]
    ],
    k = 8,
    
    return 13.

**Note:**You may assume k is always valid, 1 ≤ k ≤ n2.


## 题目大意

给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。请注意，它是排序后的第 k 小元素，而不是第 k 个元素。


说明:
你可以假设 k 的值永远是有效的, 1 ≤ k ≤ n2 。


## 解题思路


- 给出一个行有序，列有序的矩阵(并非是按照下标有序的)，要求找出这个矩阵中第 K 小的元素。注意找的第 K 小元素指的不是 k 个不同的元素，可能存在相同的元素。
- 最容易想到的就解法是优先队列。依次把矩阵中的元素推入到优先队列中。维护一个最小堆，一旦优先队列里面的元素有 k 个了，就算找到结果了。
- 这一题最优解法是二分搜索。那搜索的空间是什么呢？根据题意，可以知道，矩阵左上角的那个元素是最小的，右下角的元素是最大的。即矩阵第一个元素确定了下界，矩阵的最后一个元素确定了上界。在这个解空间里面二分搜索所有值，找到第 K  小的元素。判断是否找到的条件是，在矩阵中比 mid 小的元素个数等于 K。不断的逼近 low，使得 low == high 的时候，就是找到了第 K 小的元素了。(因为题目中说了，一定会存在第 K 小元素，所以二分搜索到一个元素的时候，一定会得出结果)。

![](https://img.halfrost.com/Leetcode/leetcode_378.png)
