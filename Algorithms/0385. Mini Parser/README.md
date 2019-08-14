# [385. Mini Parser](https://leetcode.com/problems/mini-parser/)


## 题目:

Given a nested list of integers represented as a string, implement a parser to deserialize it.

Each element is either an integer, or a list -- whose elements may also be integers or other lists.

**Note:** You may assume that the string is well-formed:

- String is non-empty.
- String does not contain white spaces.
- String contains only digits `0-9`, `[`, `-` `,`, `]`.

**Example 1:**

    Given s = "324",
    
    You should return a NestedInteger object which contains a single integer 324.

**Example 2:**

    Given s = "[123,[456,[789]]]",
    
    Return a NestedInteger object containing a nested list with 2 elements:
    
    1. An integer containing value 123.
    2. A nested list containing two elements:
        i.  An integer containing value 456.
        ii. A nested list with one element:
             a. An integer containing value 789.


## 题目大意

给定一个用字符串表示的整数的嵌套列表，实现一个解析它的语法分析器。列表中的每个元素只可能是整数或整数嵌套列表

提示：你可以假定这些字符串都是格式良好的：

- 字符串非空
- 字符串不包含空格
- 字符串只包含数字0-9, [, - ,, ]



## 解题思路

- 将一个嵌套的数据结构中的数字转换成 NestedInteger 数据结构。
- 这一题用栈一层一层的处理就行。有一些比较坑的特殊的边界数据见测试文件。这一题正确率比很多 Hard 题还要低的原因应该是没有理解好题目和边界测试数据没有考虑到。NestedInteger 这个数据结构笔者实现了一遍，见代码。
