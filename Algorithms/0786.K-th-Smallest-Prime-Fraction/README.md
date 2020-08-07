# [786. K-th Smallest Prime Fraction](https://leetcode.com/problems/k-th-smallest-prime-fraction/)


## 题目:

A sorted list `A` contains 1, plus some number of primes. Then, for every p < q in the list, we consider the fraction p/q.

What is the `K`-th smallest fraction considered? Return your answer as an array of ints, where `answer[0] = p` and `answer[1] = q`.

    Examples:
    Input: A = [1, 2, 3, 5], K = 3
    Output: [2, 5]
    Explanation:
    The fractions to be considered in sorted order are:
    1/5, 1/3, 2/5, 1/2, 3/5, 2/3.
    The third fraction is 2/5.
    
    Input: A = [1, 7], K = 1
    Output: [1, 7]

**Note:**

- `A` will have length between `2` and `2000`.
- Each `A[i]` will be between `1` and `30000`.
- `K` will be between `1` and `A.length * (A.length - 1) / 2`.


## 题目大意

一个已排序好的表 A，其包含 1 和其他一些素数.  当列表中的每一个 p<q 时，我们可以构造一个分数 p/q 。

那么第 k 个最小的分数是多少呢?  以整数数组的形式返回你的答案, 这里 answer[0] = p 且 answer[1] = q.


注意:

- A 的取值范围在 2 — 2000.
- 每个 A[i] 的值在 1 —30000.
- K 取值范围为 1 — A.length * (A.length - 1) / 2


## 解题思路


- 给出一个从小到大排列的有序数组，数组里面的元素都是质数，请找出这个数组中的数组成的真分数从小到大排列，第 K 小的分数。
- 这一题的暴力解法是枚举所有可能的真分数，从小到大排序，输出第 K 小的分数即可。注意排序的时候不能直接用 float 排序，需要转化成分子和分母的结构体进行排序。
- 最优的解法是二分搜索。由于真分数都小于 1，所以二分搜索的范围是 [0,1]。每次二分出来的 mid，需要在数组里面搜索一次，找出比 mid 小的真分数个数。并记录下最大的真分数的分子和分母，动态维护最大真分数的分子和分母。如果比 mid 小的真分数个数小于 K，那么取右区间继续二分，如果比 mid 小的真分数个数大于 K，那么取左区间继续二分。直到正好找到比 mid 小的真分数个数是 K，此时维护的最大真分数的分子和分母即为答案。
- 在已排序的矩阵中寻找最 K 小的元素这一系列的题目有：第 373 题，第 378 题，第 668 题，第 719 题，第 786 题。
