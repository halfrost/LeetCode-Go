# [925. Long Pressed Name](https://leetcode.com/problems/long-pressed-name/)

## 题目

Your friend is typing his name into a keyboard.  Sometimes, when typing a character c, the key might get long pressed, and the character will be typed 1 or more times.

You examine the typed characters of the keyboard.  Return True if it is possible that it was your friends name, with some characters (possibly none) being long pressed.



Example 1:

```c
Input: name = "alex", typed = "aaleex"
Output: true
Explanation: 'a' and 'e' in 'alex' were long pressed.
```

Example 2:

```c
Input: name = "saeed", typed = "ssaaedd"
Output: false
Explanation: 'e' must have been pressed twice, but it wasn't in the typed output.
```

Example 3:

```c
Input: name = "leelee", typed = "lleeelee"
Output: true
```

Example 4:

```c
Input: name = "laiden", typed = "laiden"
Output: true
Explanation: It's not necessary to long press any character.
```


Note:  

1. name.length <= 1000
2. typed.length <= 1000
3. The characters of name and typed are lowercase letters.

## 题目大意


给定 2 个字符串，后者的字符串中包含前者的字符串。比如在打字的过程中，某个字符会多按了几下。判断后者字符串是不是比前者字符串存在这样的“长按”键盘的情况。

## 解题思路

这一题也可以借助滑动窗口的思想。2 个字符串一起比较，如果遇到有相同的字符串，窗口继续往后滑动。直到遇到了第一个不同的字符，如果遇到两个字符串不相等的情况，可以直接返回 false。具体实现见代码。











