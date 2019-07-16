# [1002. Find Common Characters](https://leetcode.com/problems/find-common-characters/)


## 题目:

Given an array `A` of strings made only from lowercase letters, return a list of all characters that show up in all strings within the list **(including duplicates)**. For example, if a character occurs 3 times in all strings but not 4 times, you need to include that character three times in the final answer.

You may return the answer in any order.

**Example 1:**

    Input: ["bella","label","roller"]
    Output: ["e","l","l"]

**Example 2:**

    Input: ["cool","lock","cook"]
    Output: ["c","o"]

**Note:**

1. `1 <= A.length <= 100`
2. `1 <= A[i].length <= 100`
3. `A[i][j]` is a lowercase letter

## 题目大意

给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）组成的列表。例如，如果一个字符在每个字符串中出现 3 次，但不是 4 次，则需要在最终答案中包含该字符 3 次。你可以按任意顺序返回答案。


## 解题思路

- 简单题。给出一个字符串数组 A，要求找出这个数组中每个字符串都包含字符，如果字符出现多次，在最终结果中也需要出现多次。这一题可以用 map 来统计每个字符串的频次，但是如果用数组统计会更快。题目中说了只有小写字母，那么用 2 个 26 位长度的数组就可以统计出来了。遍历字符串数组的过程中，不过的缩小每个字符在每个字符串中出现的频次(因为需要找所有字符串公共的字符，公共的频次肯定就是最小的频次)，得到了最终公共字符的频次数组以后，按顺序输出就可以了。
