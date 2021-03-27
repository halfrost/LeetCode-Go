# [297. Serialize and Deserialize Binary Tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/)

## 题目

Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

**Clarification:** The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

Example 1:

```
Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]
```

Example 2:

```
Input: root = []
Output: []
```

Example 3:

```
Input: root = [1]
Output: [1]
```

Example 4:

```
Input: root = [1,2]
Output: [1,2]
```

Constraints:

* The number of nodes in the tree is in the range [0, 104].
* -1000 <= Node.val <= 1000

## 题目大意

设计一个算法，来序列化和反序列化二叉树。并不限制如何进行序列化和反序列化，但是你需要保证二叉树可以序列化为字符串，并且这个字符串可以被反序列化成原有的二叉树。

## 解题思路

1. 将给定的二叉树想象成一颗满二叉树(不存在的结点用 null 填充)。
2. 通过前序遍历，可以得到一个第一个结点为根的序列，然后递归进行序列化/反序列化即可。
