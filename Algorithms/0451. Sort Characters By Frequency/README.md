# [451. Sort Characters By Frequency](https://leetcode.com/problems/sort-characters-by-frequency/)

## 题目

Given a string, sort it in decreasing order based on the frequency of characters.


Example 1:

```c
Input:
"tree"

Output:
"eert"

Explanation:
'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
```

Example 2:

```c
Input:
"cccaaa"

Output:
"cccaaa"

Explanation:
Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
Note that "cacaca" is incorrect, as the same characters must be together.
```

Example 3:

```c
Input:
"Aabb"

Output:
"bbAa"

Explanation:
"bbaA" is also a valid answer, but "Aabb" is incorrect.
Note that 'A' and 'a' are treated as two different characters.
```




## 题目大意

这道题是 Google 的面试题。

给定一个字符串，要求根据字符出现的频次从高到低重新排列这个字符串。

## 解题思路

思路比较简单，首先统计每个字符的频次，然后排序，最后按照频次从高到低进行输出即可。



