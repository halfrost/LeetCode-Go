# [208. Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/)


## 题目

Implement a trie with `insert`, `search`, and `startsWith` methods.

**Example:**

    Trie trie = new Trie();
    
    trie.insert("apple");
    trie.search("apple");   // returns true
    trie.search("app");     // returns false
    trie.startsWith("app"); // returns true
    trie.insert("app");   
    trie.search("app");     // returns true

**Note:**

- You may assume that all inputs are consist of lowercase letters `a-z`.
- All inputs are guaranteed to be non-empty strings.

## 题目大意

实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

## 解题思路

- 要求实现一个 Trie 的数据结构，具有 `insert`, `search`, `startsWith` 三种操作
- 这一题就是经典的 Trie 实现。本题的实现可以作为 Trie 的模板。
