# [354. Russian Doll Envelopes](https://leetcode.com/problems/russian-doll-envelopes/)


## 题目:

You have a number of envelopes with widths and heights given as a pair of integers `(w, h)`. One envelope can fit into another if and only if both the width and height of one envelope is greater than the width and height of the other envelope.

What is the maximum number of envelopes can you Russian doll? (put one inside other)

**Note:**Rotation is not allowed.

**Example:**

    **Input:** [[5,4],[6,4],[6,7],[2,3]]
    **Output:** 3
    **Explanation: T**he maximum number of envelopes you can Russian doll is 3 ([2,3] => [5,4] => [6,7]).


## 题目大意

给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h) 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

说明:
- 不允许旋转信封。

## 解题思路

- 给出一组信封的宽度和高度，如果组成俄罗斯套娃，问最多能套几层。只有当一个信封的宽度和高度都比另外一个信封大的时候，才能套在小信封上面。
- 这一题的实质是第 300 题 Longest Increasing Subsequence 的加强版。能组成俄罗斯套娃的条件就是能找到一个最长上升子序列。但是这题的条件是二维的，要求能找到在二维上都能满足条件的最长上升子序列。先降维，把宽度排序。然后在高度上寻找最长上升子序列。这里用到的方法和第 300 题的方法一致。解题思路详解见第 300 题。
