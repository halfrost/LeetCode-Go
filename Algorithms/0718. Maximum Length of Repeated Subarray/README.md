# [718. Maximum Length of Repeated Subarray](https://leetcode.com/problems/maximum-length-of-repeated-subarray/)


## 题目:

Given two integer arrays `A` and `B`, return the maximum length of an subarray that appears in both arrays.

**Example 1:**

    Input:
    A: [1,2,3,2,1]
    B: [3,2,1,4,7]
    Output: 3
    Explanation: 
    The repeated subarray with maximum length is [3, 2, 1].

**Note:**

1. 1 <= len(A), len(B) <= 1000
2. 0 <= A[i], B[i] < 100


## 题目大意

给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。



## 解题思路

- 给出两个数组，求这两个数组中最长相同子串的长度。
- 这一题最容易想到的是 DP 动态规划的解法。`dp[i][j]` 代表在 A 数组中以 `i` 下标开始的子串与 B 数组中以 `j` 下标开始的子串最长相同子串的长度，状态转移方程为 `dp[i][j] = dp[i+1][j+1] + 1` (当 `A[i] == B[j]`)。这种解法的时间复杂度是 O(n^2)，空间复杂度 O(n^2)。
- 这一题最佳解法是二分搜索 + `Rabin-Karp`。比较相同子串耗时的地方在于，需要一层循环，遍历子串所有字符。但是如果比较两个数字就很快，`O(1)` 的时间复杂度。所以有人就想到了，能不能把字符串也映射成数字呢？这样比较起来就非常快。这个算法就是 `Rabin-Karp` 算法。字符串映射成一个数字不能随意映射，还要求能根据字符串前缀动态增加，比较下一个字符串的时候，可以利用已比较过的前缀，加速之后的字符串比较。在 Rabin-Karp 算法中有一个“码点”的概念。类似于10进制中的进制。具体的算法讲解可以见这篇：

    [基础知识 - Rabin-Karp 算法](https://www.cnblogs.com/golove/p/3234673.html)

    “码点”一般取值为一个素数。在 go 的 `strings` 包里面取值是 16777619。所以这一题也可以直接取这个值。由于这一次要求我们找最长长度，所以把最长长度作为二分搜索的目标。先将数组 A 和数组 B 中的数字都按照二分出来的长度，进行 `Rabin-Karp` hash。对 A 中的 hash 与下标做映射关系，存到 map 中，方便后面快速查找。然后遍历 B 中的 hash，当 hash 一致的时候，再匹配下标。如果下标存在，且拥有相同的前缀，那么就算找到了相同的子串了。最后就是不断的二分，找到最长的结果即可。这个解法的时间复杂度 O(n * log n)，空间复杂度 O(n)。