# [745. Prefix and Suffix Search](https://leetcode.com/problems/prefix-and-suffix-search/)


## 题目

Given many `words`, `words[i]` has weight `i`.

Design a class `WordFilter` that supports one function, `WordFilter.f(String prefix, String suffix)`. It will return the word with given `prefix` and `suffix` with maximum weight. If no word exists, return -1.

**Examples:**

    Input:
    WordFilter(["apple"])
    WordFilter.f("a", "e") // returns 0
    WordFilter.f("b", "") // returns -1

**Note:**

1. `words` has length in range `[1, 15000]`.
2. For each test case, up to `words.length` queries `WordFilter.f` may be made.
3. `words[i]` has length in range `[1, 10]`.
4. `prefix, suffix` have lengths in range `[0, 10]`.
5. `words[i]` and `prefix, suffix` queries consist of lowercase letters only.


## 题目大意

给定多个 words，words[i] 的权重为 i 。设计一个类 WordFilter 实现函数WordFilter.f(String prefix, String suffix)。这个函数将返回具有前缀 prefix 和后缀suffix 的词的最大权重。如果没有这样的词，返回 -1。



## 解题思路


- 要求实现一个 `WordFilter` ，它具有字符串匹配的功能，可以匹配出前缀和后缀都满足条件的字符串下标，如果找得到，返回下标，如果找不到，则返回 -1 。
- 这一题有 2 种解题思路。第一种是先把这个 `WordFilter` 结构里面的字符串全部预处理一遍，将它的前缀，后缀的所有组合都枚举出来放在 map 中，之后匹配的时候只需要按照自己定义的规则查找 key 就可以了。初始化时间复杂度 `O(N * L^2)`，查找时间复杂度 `O(1)`，空间复杂度 `O(N * L^2)`。其中 `N` 是输入的字符串数组的长度，`L` 是输入字符串数组中字符串的最大长度。第二种思路是直接遍历字符串每个下标，依次用字符串的前缀匹配方法和后缀匹配方法，依次匹配。初始化时间复杂度 `O(1)`，查找时间复杂度 `O(N * L)`，空间复杂度 `O(1)`。其中 `N` 是输入的字符串数组的长度，`L` 是输入字符串数组中字符串的最大长度。

