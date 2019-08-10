# [433. Minimum Genetic Mutation](https://leetcode.com/problems/minimum-genetic-mutation/)


## 题目

A gene string can be represented by an 8-character long string, with choices from `"A"`, `"C"`, `"G"`, `"T"`.

Suppose we need to investigate about a mutation (mutation from "start" to "end"), where ONE mutation is defined as ONE single character changed in the gene string.

For example, `"AACCGGTT"` -> `"AACCGGTA"` is 1 mutation.

Also, there is a given gene "bank", which records all the valid gene mutations. A gene must be in the bank to make it a valid gene string.

Now, given 3 things - start, end, bank, your task is to determine what is the minimum number of mutations needed to mutate from "start" to "end". If there is no such a mutation, return -1.

**Note:**

1. Starting point is assumed to be valid, so it might not be included in the bank.
2. If multiple mutations are needed, all mutations during in the sequence must be valid.
3. You may assume start and end string is not the same.

**Example 1:**

    start: "AACCGGTT"
    end:   "AACCGGTA"
    bank: ["AACCGGTA"]
    
    return: 1

**Example 2:**

    start: "AACCGGTT"
    end:   "AAACGGTA"
    bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
    
    return: 2

**Example 3:**

    start: "AAAAACCC"
    end:   "AACCCCCC"
    bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
    
    return: 3


## 题目大意

现在给定3个参数 — start, end, bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变化次数。如果无法实现目标变化，请返回 -1。

注意:

1. 起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
2. 所有的目标基因序列必须是合法的。
3. 假定起始基因序列与目标基因序列是不一样的。


## 解题思路


- 给出 start 和 end 两个字符串和一个 bank 字符串数组，问从 start 字符串经过多少次最少变换能变换成 end 字符串。每次变换必须使用 bank 字符串数组中的值。
- 这一题完全就是第 127 题的翻版题，解题思路和代码 99% 是一样的。相似的题目也包括第 126 题。这一题比他们都要简单。有 2 种解法，BFS 和 DFS。具体思路可以见第 127 题的题解。

