# [76. Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/)

## 题目

Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

Example:

```c
Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
```

Note:    

- If there is no such window in S that covers all characters in T, return the empty string "".
- If there is such window, you are guaranteed that there will always be only one unique minimum window in S.

## 题目大意

给定一个源字符串 s，再给一个字符串 T，要求在源字符串中找到一个窗口，这个窗口包含由字符串各种排列组合组成的，窗口中可以包含 T 中没有的字符，如果存在多个，在结果中输出最小的窗口，如果找不到这样的窗口，输出空字符串。

## 解题思路

这一题是滑动窗口的题目，在窗口滑动的过程中不断的包含字符串 T，直到完全包含字符串 T 的字符以后，记下左右窗口的位置和窗口大小。每次都不断更新这个符合条件的窗口和窗口大小的最小值。最后输出结果即可。




