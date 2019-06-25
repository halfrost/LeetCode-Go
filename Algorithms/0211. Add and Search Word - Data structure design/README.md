# [211. Add and Search Word - Data structure design](https://leetcode.com/problems/add-and-search-word-data-structure-design/)


## 题目

Design a data structure that supports the following two operations:

    void addWord(word)
    bool search(word)

search(word) can search a literal word or a regular expression string containing only letters `a-z` or `.`. A `.` means it can represent any one letter.

**Example:**

    addWord("bad")
    addWord("dad")
    addWord("mad")
    search("pad") -> false
    search("bad") -> true
    search(".ad") -> true
    search("b..") -> true

**Note:**You may assume that all words are consist of lowercase letters `a-z`.

## 题目大意

设计一个支持以下两种操作的数据结构：`void addWord(word)`、`bool search(word)`。`search(word)` 可以搜索文字或正则表达式字符串，字符串只包含字母 . 或 a-z 。 "." 可以表示任何一个字母。



## 解题思路

- 设计一个 `WordDictionary` 的数据结构，要求具有 `addWord(word)` 和 `search(word)` 的操作，并且具有模糊查找的功能。
- 这一题是第 208 题的加强版，在第 208 题经典的 Trie 上加上了模糊查找的功能。其他实现一模一样。
