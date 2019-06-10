# [1054. Distant Barcodes](https://leetcode.com/problems/distant-barcodes/)

## 题目

In a warehouse, there is a row of barcodes, where the i-th barcode is barcodes[i].

Rearrange the barcodes so that no two adjacent barcodes are equal.  You may return any answer, and it is guaranteed an answer exists.

 

Example 1:

```c
Input: [1,1,1,2,2,2]
Output: [2,1,2,1,2,1]
```

Example 2:

```c
Input: [1,1,1,1,2,2,3,3]
Output: [1,3,1,3,2,1,2,1]
```
 

Note:

1. 1 <= barcodes.length <= 10000
2. 1 <= barcodes[i] <= 10000


## 题目大意

给出一个条形二维码代码数组，要求重新排列这组数字，保证相邻的两两数字不相等。输出任意一个解即可。题目保证能有一组解。


## 解题思路

- 这一题和第 767 题原理是完全一样的。第 767 题是 Google 的面试题。
- 解题思路比较简单，先按照每个数字的频次从高到低进行排序，注意会有频次相同的数字。排序以后，分别从第 0 号位和中间的位置开始往后取数，取完以后即为最终解。