# [567. Permutation in String](https://leetcode.com/problems/permutation-in-string/)

## 题目

Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1. In other words, one of the first string's permutations is the substring of the second string.


Example 1:

```c
Input:s1 = "ab" s2 = "eidbaooo"
Output:True
Explanation: s2 contains one permutation of s1 ("ba").
```

Example 2:

```c
Input:s1= "ab" s2 = "eidboaoo"
Output: False
```

Note:

1. The input strings only contain lower case letters.
2. The length of both given strings is in range [1, 10,000].

## 题目大意


在一个字符串重寻找子串出现的位置。子串可以是 Anagrams 形式存在的。Anagrams 是一个字符串任意字符的全排列组合。

## 解题思路

这一题和第 438 题，第 3 题，第 76 题，第 567 题类似，用的思想都是"滑动窗口"。


这道题只需要判断是否存在，而不需要输出子串所在的下标起始位置。所以这道题是第 438 题的缩水版。具体解题思路见第 438 题。






